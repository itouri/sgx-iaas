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
//#include <openssl/evp.h>
#include "../message.h"

#include "e_fileio.h"

// #include "stdlib.h"
// #include "string.h"
// #include "sgx_tcrypto.h"
// #include "se_tcrypto_common.h"
// #include "openssl/aes.h"
// #include "openssl/evp.h"

#define MAC_SIZE 16

/**
 * Encryption
 */
int
elgamal_encrypt(byte **encData, byte *data, int dataLen, const EC_KEY *eckey) 
{
	BN_CTX *ctx = NULL;
	BIGNUM *r = NULL, *p = NULL, *m;
	EC_POINT *C1 = NULL, *C2 = NULL;
	EC_POINT *Tmp = NULL, *M;
	const EC_POINT *Pkey;
	const EC_GROUP *group;
	int    c1Len, c2Len;
	int    rv;

	if ((group = EC_KEY_get0_group(eckey)) == NULL) {
		return 0;
	}
	p = BN_new();
	ctx = BN_CTX_new();
	EC_GROUP_get_curve_GFp(group, p, NULL, NULL, ctx);

	// printf(" p = ");
	// BN_print_fp(stdout, p);
	// puts("");

	/* C1 = r*G */
	C1 = EC_POINT_new(group);

	/* generate random number r */ 
	r = BN_new();
	M = EC_POINT_new(group);
	m = BN_new();
	do {
		if (!BN_rand_range(r, p)) {
			return 0;
		}
	} while (BN_is_zero(r));
	// printf(" r = ");
	// BN_print_fp(stdout, r);
	// puts("");

	EC_POINT_mul(group, C1, r, NULL, NULL, ctx);

	/* C2 = r*P + M */ 
	/* M */
	BN_bin2bn(data, dataLen, m);
	rv = EC_POINT_set_compressed_coordinates_GFp(group, M, m, 1, ctx);
	if (!rv) {
		printf("error EC_POINT_set_compressed_coordinates_GFp");
		return 0;
	}

	C2 = EC_POINT_new(group);
	Tmp = EC_POINT_new(group);
	Pkey = EC_KEY_get0_public_key(eckey);
	EC_POINT_mul(group, Tmp, NULL, Pkey, r, ctx);
	EC_POINT_add(group, C2, Tmp, M, ctx);

	/* cipher text C = (C1, C2) */ 
	c1Len = EC_POINT_point2oct(group, C1, POINT_CONVERSION_COMPRESSED,
							   NULL, 0, ctx);
	printf(" Point converted length (C1) = %d\n", c1Len);
	c2Len =	EC_POINT_point2oct(group, C2, POINT_CONVERSION_COMPRESSED,
							   NULL, 0, ctx);
	printf(" Point converted length (C2) = %d\n", c1Len);
	*encData = OPENSSL_malloc(c1Len + c2Len);
	EC_POINT_point2oct(group, C1, POINT_CONVERSION_COMPRESSED,
							*encData, c1Len, ctx);
	EC_POINT_point2oct(group, C2, POINT_CONVERSION_COMPRESSED,
							*encData + c1Len, c2Len, ctx);

	BN_clear_free(p);
	BN_clear_free(r);
	BN_clear_free(m);
	EC_POINT_free(C1);
	EC_POINT_free(C2);
	EC_POINT_free(M);
	EC_POINT_free(Tmp);
	BN_CTX_free(ctx);

	return (c1Len + c2Len);
}

/**
 * Decryption
 */
int
elgamal_decrypt(byte **decData, byte *encData, int encLen, const EC_KEY *eckey) 
{
	int rv;
	const EC_GROUP *group;
	const BIGNUM *prvKey;
	BN_CTX *ctx;
	EC_POINT *C1 = NULL, *C2 = NULL;
	EC_POINT *M = NULL, *Tmp = NULL;

	group = EC_KEY_get0_group(eckey);
	prvKey = EC_KEY_get0_private_key(eckey);
#ifdef DEBUG
	printf(" prvKey = ");
	BN_print_fp(stdout, prvKey);
	puts("");
#endif
	C1 = EC_POINT_new(group);
	C2 = EC_POINT_new(group);
	ctx = BN_CTX_new();

	/* C1 */
#ifdef DEBUG
	printHex("C1", encData, encLen / 2);
#endif
	rv = EC_POINT_oct2point(group, C1, encData, encLen / 2, ctx);
	if (!rv) {
		printf("EC_POINT_oct2point error (C1)\n");
		return 0;
	}

	/* C2 */
#ifdef DEBUG
	printHex("C2", encData + encLen / 2, encLen / 2);
#endif
	rv = EC_POINT_oct2point(group, C2, encData + encLen / 2, encLen / 2,
							ctx);
	if (!rv) {
		printf("EC_POINT_oct2point error (C2)\n");
		return 0;
	}
	Tmp = EC_POINT_new(group);
	M = EC_POINT_new(group);

	/* M = C2 - x C1 */ 
	EC_POINT_mul(group, Tmp, NULL, C1, prvKey, ctx);
	EC_POINT_invert(group, Tmp, ctx);
	EC_POINT_add(group, M, C2, Tmp, ctx);

	/* Output M */ 
	rv = EC_POINT_point2oct(group, M, POINT_CONVERSION_COMPRESSED, NULL, 0,
							ctx);

#ifdef DEBUG
	printf(" Point converted length = %d\n", rv);
#endif
	*decData = OPENSSL_malloc(rv);
	EC_POINT_point2oct(group, M, POINT_CONVERSION_COMPRESSED, *decData,
					   rv, ctx);

	EC_POINT_free(C1);
	EC_POINT_free(C2);
	EC_POINT_free(M);
	EC_POINT_free(Tmp);
	BN_CTX_free(ctx);

	return rv;
}
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

	sha_ret= sgx_sha256_msg((const uint8_t *) &ra_key, sizeof(ra_key), 
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
sgx_status_t enclave_launch_vm(unsigned char *cry_req_data, uuid_t *image_id, sgx_ra_context_t *ctx)
{
	sgx_status_t ret;
	int retval;
	// MRSINGERはいらないかな
	//Task image復号化鍵を受け取る
	char cry_msg[512];
	recv_from_ras((char **)cry_msg, (size_t*)sizeof(cry_msg));

	// ra共通鍵でimage復号鍵を復号化
	uint8_t aes_gcm_iv[12] = {0};
	//TODO 鍵がただの文字列なのは良くない
	unsigned char imd_key[16]; //EVP_PKEY *imd_key;
	sgx_aes_gcm_128bit_tag_t mac[16];
	ret = sgx_rijndael128GCM_decrypt(
		&ra_key,
		(const uint8_t*)cry_msg,
		sizeof(imd_key),
		imd_key,
		&aes_gcm_iv[0],
		12,
		NULL,
		0,
		mac //TODO このmacって検証しなくていいの？関数がしてくれないの？
	);

	

	const char iv = {0}; // 16Byte
	unsigned char *vrfymac;
	// cry_client_id を復号 -> cry_req_data を復号
	req_data_t *req_data;
	decrypt((const char*)imd_key, (const unsigned char*)req_data, sizeof(req_data), iv, (unsigned char*)req_data, sizeof(req_data_t));

	// TODO macを検証
	// cmac128(imd_key, (unsigned char *) req_data,
	// 	sizeof(req_data_t),
	// 	(unsigned char *) vrfymac);
	
	// if ( CRYPTO_memcmp(req_data->mac, vrfymac, sizeof(MAC_SIZE)) ) {
	// 	printe("Failed to verify request data MAC\n");
	// 	return 0;
	// }
	// ret = sgx_rijndael128_cmac_msg(
	// 	imd_key,
	// 	(const uint8_t*) req_data,
	// 	sizeof(req_data),
	// 	p_mac
	// );

	// grapheneSGX の mrenclave を RAS に問い合わせる
	sgx_measurement_t *vm_mrenclave;
	recv_from_ras((char**)cry_msg, (size_t*)sizeof(cry_msg));

	ret = sgx_rijndael128GCM_decrypt(
		&ra_key,
		(const uint8_t*)cry_msg,
		sizeof(cry_msg), //TODO sizeが間違ってる
		(uint8_t*)vm_mrenclave,
		&aes_gcm_iv[0],
		12,
		NULL,
		0,
		mac //TODO このmacって検証しなくていいの？関数がしてくれないの？
	);

	//Task 指定された graphene で image を起動
	sgx_enclave_id_t graphene_eid; // graphene からの id を返却してもらう
	//TODO (unsigned char (*)[16])大丈夫かなこのキャスト
	run_graphene_vm_ocall(&retval, &graphene_eid, (unsigned char (*)[16])image_id);
	// if (ret != 0) {
	// 	return Error;
	// }

	//Task この enclave が まず通常の LA
	// やること多いぞ

	//Task OKなら imageID, clientID を RAS へ報告
	// image_id, client_id がどちらもuuidにすれば固定長にできる筈
	// *image_idの * いる？
	msg_cmpt_t msg_cmpt;
	memcpy(&msg_cmpt.image_id, &image_id, sizeof(uuid_t));
	memcpy(&msg_cmpt.client_id, &req_data->client_id, sizeof(uuid_t));

	uint8_t crypted_msg[512];
	ret = sgx_rijndael128GCM_encrypt(
		&ra_key,
		(const uint8_t*)&msg_cmpt,
		sizeof(msg_cmpt_t),
		crypted_msg,
		&aes_gcm_iv[0],
		12,
		NULL,
		0,
		mac
	);
	send_to_ras((char*)crypted_msg, sizeof(crypted_msg));
}
