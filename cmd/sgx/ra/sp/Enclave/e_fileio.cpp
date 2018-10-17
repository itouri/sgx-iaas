#include <stdio.h>
#include <sys/types.h>

#include "e_fileio.h"
#include "Enclave_t.h"

// RSAへのsend
int send_to_ras(char *src, size_t sz) {
	int ret_val;
	send_to_ras_ocall(&ret_val, src ,sz);
	return ret_val;
}

// RSAからのreceive
int recv_from_ras(char **dest, size_t *sz) {
	int ret_val;
	recv_from_ras_ocall(&ret_val, dest, sz);
	return ret_val;
}

// localfileのread
int read_file(unsigned char *dest, char *file, off_t *len) {
	int ret_val;
	read_file_ocall(&ret_val, dest, file, len);
	return ret_val;
}