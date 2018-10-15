/*

Copyright 2018 Intel Corporation

This software and the related documents are Intel copyrighted materials,
and your use of them is governed by the express license under which they
were provided to you (License). Unless the License provides otherwise,
you may not use, modify, copy, publish, distribute, disclose or transmit
this software or the related documents without Intel's prior written
permission.

This software and the related documents are provided as is, with no
express or implied warranties, other than those that are expressly stated
in the License.

*/

#ifndef _WIN32
#include "../config.h"
#endif
#include "Enclave_t.h"
#include <string.h>
#include <sgx_utils.h>
#include <sgx_tae_service.h>
#include <sgx_tkey_exchange.h>
#include <sgx_tcrypto.h>

// 追加
#include <sgx_report.h>
#include <uuid/uuid.h>
#include <openssl/evp.h>
#include "../message.h"

#include "e_fileio.h"

#define MAC_SIZE 16


static const sgx_ec256_public_t def_service_public_key = {
    {
        0x72, 0x12, 0x8a, 0x7a, 0x17, 0x52, 0x6e, 0xbf,
        0x85, 0xd0, 0x3a, 0x62, 0x37, 0x30, 0xae, 0xad,
        0x3e, 0x3d, 0xaa, 0xee, 0x9c, 0x60, 0x73, 0x1d,
        0xb0, 0x5b, 0xe8, 0x62, 0x1c, 0x4b, 0xeb, 0x38
    },
    {
        0xd4, 0x81, 0x40, 0xd9, 0x50, 0xe2, 0x57, 0x7b,
        0x26, 0xee, 0xb7, 0x41, 0xe7, 0xc6, 0x14, 0xe2,
        0x24, 0xb7, 0xbd, 0xc9, 0x03, 0xf2, 0x9a, 0x28,
        0xa8, 0x3c, 0xc8, 0x10, 0x11, 0x14, 0x5e, 0x06
    }

};

#define PSE_RETRIES	5	/* Arbitrary. Not too long, not too short. */

/*----------------------------------------------------------------------
 * WARNING
 *----------------------------------------------------------------------
 *
 * End developers should not normally be calling these functions
 * directly when doing remote attestation:
 *
 *    sgx_get_ps_sec_prop()
 *    sgx_get_quote()
 *    sgx_get_quote_size()
 *    sgx_get_report()
 *    sgx_init_quote()
 *
 * These functions short-circuits the RA process in order
 * to generate an enclave quote directly!
 *
 * The high-level functions provided for remote attestation take
 * care of the low-level details of quote generation for you:
 *
 *   sgx_ra_init()
 *   sgx_ra_get_msg1
 *   sgx_ra_proc_msg2
 *
 *----------------------------------------------------------------------
 */

/*
 * This doesn't really need to be a C++ source file, but a bug in 
 * 2.1.3 and earlier implementations of the SGX SDK left a stray
 * C++ symbol in libsgx_tkey_exchange.so so it won't link without
 * a C++ compiler. Just making the source C++ was the easiest way
 * to deal with that.
 */

sgx_status_t get_report(sgx_report_t *report, sgx_target_info_t *target_info)
{
#ifdef SGX_HW_SIM
	return sgx_create_report(NULL, NULL, report);
#else
	return sgx_create_report(target_info, NULL, report);
#endif
}

size_t get_pse_manifest_size ()
{
	return sizeof(sgx_ps_sec_prop_desc_t);
}

sgx_status_t get_pse_manifest(char *buf, size_t sz)
{
	sgx_ps_sec_prop_desc_t ps_sec_prop_desc;
	sgx_status_t status= SGX_ERROR_SERVICE_UNAVAILABLE;
	int retries= PSE_RETRIES;

	do {
		status= sgx_create_pse_session();
		if ( status != SGX_SUCCESS ) return status;
	} while (status == SGX_ERROR_BUSY && retries--);
	if ( status != SGX_SUCCESS ) return status;

	status= sgx_get_ps_sec_prop(&ps_sec_prop_desc);
	if ( status != SGX_SUCCESS ) return status;

	memcpy(buf, &ps_sec_prop_desc, sizeof(ps_sec_prop_desc));

	sgx_close_pse_session();

	return status;
}

sgx_status_t enclave_ra_init(sgx_ec256_public_t key, int b_pse,
	sgx_ra_context_t *ctx, sgx_status_t *pse_status)
{
	sgx_status_t ra_status;

	/*
	 * If we want platform services, we must create a PSE session 
	 * before calling sgx_ra_init()
	 */

	if ( b_pse ) {
		int retries= PSE_RETRIES;
		do {
			*pse_status= sgx_create_pse_session();
			if ( *pse_status != SGX_SUCCESS ) return SGX_ERROR_UNEXPECTED;
		} while (*pse_status == SGX_ERROR_BUSY && retries--);
		if ( *pse_status != SGX_SUCCESS ) return SGX_ERROR_UNEXPECTED;
	}

	ra_status= sgx_ra_init(&key, b_pse, ctx);

	if ( b_pse ) {
		int retries= PSE_RETRIES;
		do {
			*pse_status= sgx_create_pse_session();
			if ( *pse_status != SGX_SUCCESS ) return SGX_ERROR_UNEXPECTED;
		} while (*pse_status == SGX_ERROR_BUSY && retries--);
		if ( *pse_status != SGX_SUCCESS ) return SGX_ERROR_UNEXPECTED;
	}

	return ra_status;
}

sgx_status_t enclave_ra_init_def(int b_pse, sgx_ra_context_t *ctx,
	sgx_status_t *pse_status)
{
	return enclave_ra_init(def_service_public_key, b_pse, ctx, pse_status);
}

/*
 * Return a SHA256 hash of the requested key. KEYS SHOULD NEVER BE
 * SENT OUTSIDE THE ENCLAVE IN PLAIN TEXT. This function let's us
 * get proof of possession of the key without exposing it to untrusted
 * memory.
 */

//TODO 大丈夫これ？
sgx_ra_key_128_t ra_key;

// ここに入る段階では enclaveは信頼できていて RAS との安全な通信が確保されている
sgx_status_t enclave_ra_get_key_hash(sgx_status_t *get_keys_ret,
	sgx_ra_context_t ctx, sgx_ra_key_type_t type, sgx_sha256_hash_t *hash)
{
	sgx_status_t sha_ret;
	//sgx_ra_key_128_t k;

	// First get the requested key which is one of:
	//  * SGX_RA_KEY_MK 
	//  * SGX_RA_KEY_SK
	// per sgx_ra_get_keys().

	//*get_keys_ret= sgx_ra_get_keys(ctx, type, &k);
	*get_keys_ret= sgx_ra_get_keys(ctx, type, &ra_key);
	if ( *get_keys_ret != SGX_SUCCESS ) return *get_keys_ret;

	/* Now generate a SHA hash */

	sha_ret= sgx_sha256_msg((const uint8_t *) &k, sizeof(k), 
		(sgx_sha256_hash_t *) hash); // Sigh.

	/* Let's be thorough */

	// この k がサーバーと交換した鍵？ enclave_launch_vmにわたす 
	//memset(k, 0, sizeof(k));

	return sha_ret;
}

sgx_status_t enclave_ra_close(sgx_ra_context_t ctx)
{
        sgx_status_t ret;
        ret = sgx_ra_close(ctx);
        return ret;
}

// manifest復号鍵を基にmanifestを復号化してVMを起動する
sgx_status_t enclave_launch_vm(unsigned char *cry_req_data, uuid_t image_id, sgx_ra_context_t ctx)
{
	// MRSINGERはいらないかな
	//Task image復号化鍵を受け取る
	unsigned char *cry_msg;
	recv_from_ras(cry_msg, sizeof(cry_msg));

	// ra共通鍵でimage復号鍵を復号化
	uint8_t aes_gcm_iv[12] = {0};
	//TODO 鍵がただの文字列なのは良くない
	unsigned char imd_key[16]; //EVP_PKEY *imd_key;
	sgx_aes_gcm_128bit_tag_t mac[16];
	ret = sgx_rijndael128GCM_decrypt(
		&ra_key,
		cry_msg,
		sizeof(imd_key),
		imd_key,
		&aes_gcm_iv[0],
		12,
		NULL,
		0,
		mac //TODO このmacって検証しなくていいの？関数がしてくれないの？
	);

	// --- 仕様変更で image_meta_data を読み込む必要はなくなった！！！ ---
	// // image metadata を読み込む
	// //固定長にする必要がある
	// unsigned char crypted_imd[32];
	// char uuid[36+1];
    // uuid_unparse(image_id, uuid);
	// char* image_path = IMAGE_PATH + uuid;
	// read_file(image_path, crypted_imd);

	// // image復号鍵でimage_metadataを復号化
	// // sgx_rijndael128GCM_decrypt ではダメ？
	// unsigned char iv = {0}; // 16Byte
	// image_metadata_t *image_metadata;
	// decrypt(imd_key, crypted_imd, sizeof(crypted_imd), iv, image_metadata, sizeof(image_meta_data_t));

	// //Task MRENCLAVE, clientIDに改ざんがないかMACで検証	
	// //TODO １つめの鍵は何を指定しよう
	// cmac128(imd_key, (unsigned char *) image_metadata,
	// 	sizeof(image_metadata_t),
	// 	(unsigned char *) vrfymac);
	// if ( CRYPTO_memcmp(image_metadata.mac, vrfymac, sizeof(MAC_SIZE)) ) {
	// 	eprintf("Failed to verify image_metadata MAC\n");
	// 	return 0;
	// }
	// --- 仕様変更で image_meta_data を読み込む必要はなくなった！！！ ---

	// cry_client_id を復号 -> cry_req_data を復号
	req_data_t *req_data;
	decrypt(imd_key, req_data, sizeof(req_data), iv, req_data, sizeof(req_data_t));

	// macを検証
	cmac128(imd_key, (unsigned char *) req_data,
		sizeof(req_data_t),
		(unsigned char *) vrfymac);
	if ( CRYPTO_memcmp(req_data.mac, vrfymac, sizeof(MAC_SIZE)) ) {
		eprintf("Failed to verify request data MAC\n");
		return 0;
	}

	// grapheneSGX の mrenclave を RAS に問い合わせる
	sgx_measurement_t *vm_mrenclave
	recv_from_ras(cry_msg, sizeof(cry_msg));

	ret = sgx_rijndael128GCM_decrypt(
		&ra_key,
		cry_msg,
		sizeof(cry_msg), //TODO sizeが間違ってる
		vm_mrenclave,
		&aes_gcm_iv[0],
		12,
		NULL,
		0,
		mac //TODO このmacって検証しなくていいの？関数がしてくれないの？
	);

	//Task 指定された graphene で image を起動
	sgx_enclave_id_t graphene_eid; // graphene からの id を返却してもらう
	ret = run_graphene_vm_ocall(&graphene_eid, image_id);
	if (ret != 0) {
		return;
	}

	//Task この enclave が まず通常の LA
	// やること多いぞ

	//Task OKなら imageID, clientID を RAS へ報告
	// image_id, client_id がどちらもuuidにすれば固定長にできる筈
	msg_cmpt_t msg_cmpt = {image_id, cliend_id};
	uint8_t crypted_msg[];
	ret = sgx_rijndael128GCM_encrypt(
		&ra_key,
		msg_cmpt,
		sizeof(msg_cmpt_t),
		crypted_msg,
		&aes_gcm_iv[0],
		12,
		NULL,
		0,
		mac
	);
	send_to_ras(crypted_msg, sizeof(sz));
}


// https://blanktar.jp/blog/2014/10/c_language-aes-with-openssl.html
unsigned char* decrypt(
    const char* key,
    const unsigned char* data,
    const size_t datalen,
    const unsigned char* iv,
    unsigned char* dest,
    const size_t destlen
){
	EVP_CIPHER_CTX *de;
    // え？enclaveでなんで malloc できるん？
    de = (EVP_CIPHER_CTX *)malloc(sizeof(EVP_CIPHER_CTX));

	int f_len = 0;
	int p_len = datalen;

	memset(dest, 0x00, destlen);

	EVP_CIPHER_CTX_init(de);
	EVP_DecryptInit_ex(de, EVP_aes_128_cbc(), NULL, (unsigned char*)key, iv);

	EVP_DecryptUpdate(de, (unsigned char *)dest, &p_len, data, datalen);
	//EVP_DecryptFinal_ex(&de, (unsigned char *)(dest + p_len), &f_len);

	EVP_CIPHER_CTX_cleanup(de);

	return dest;
}

// MRENCLAVE を rsa 公開鍵で暗号化するために使う
unsigned char* encrypt( const char* key,
						const unsigned char* data,
						const size_t datalen,
						const unsigned char* iv,
						unsigned char* dest,
						const size_t destlen)
{
	EVP_CIPHER_CTX *en;
    // やっぱmallocが必要だよね どうやら malloc をラップしているライブラリがあるみたい
    en = (EVP_CIPHER_CTX *)malloc(sizeof(EVP_CIPHER_CTX));
    EVP_CIPHER_CTX_init(en);

	int i, f_len=0;
	int c_len = destlen;

	memset(dest, 0x00, destlen);

	EVP_EncryptInit_ex(en, EVP_aes_128_cbc(), NULL, (unsigned char*)key, iv);

	EVP_EncryptUpdate(en, dest, &c_len, (unsigned char *)data, datalen);
	//EVP_EncryptFinal_ex(&en, (unsigned char *)(dest + c_len), &f_len);

	EVP_CIPHER_CTX_cleanup(en);

	return dest;
}

