#include "Enclave_t.h"

// RSAへのsend
int send_to_ras(void *src, size_t sz) {
	return send_to_ras_ocall(src ,sz);
}

// RSAからのreceive
int recv_from_ras(void **dest, size_t *sz) {
	return recv_from_ras_ocall(dest, sz);
}

// localfileのread
int read_file(unsigned char *dest, char *file, off_t *len) {
	return read_file_ocall(dest, file, len);
}