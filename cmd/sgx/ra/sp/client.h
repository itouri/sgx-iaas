#ifndef __CLIENT_H
#define __CLIENT_H

#include <sgx_utils.h>
#include <stdio.h>
#include "uuid/uuid.h"

int run_graphene_vm_ocall(sgx_enclave_id_t *graphene_eid, uuid_t image_id);

int send_to_ras_ocall(void *src, size_t sz);

// RSAからのreceive
int recv_from_ras_ocall(void **dest, size_t *sz);

// localfileのread
int read_file_ocall(unsigned char *dest, char *file, off_t *len);

#endif