#ifndef __E_FILEIO_H
#define __E_FILEIO_H

#include <sgx_utils.h>
#include "Enclave_t.h"

int send_to_ras(void *src, size_t sz);
int recv_from_ras(void **dest, size_t *sz);
int read_file(unsigned char *dest, char *file, off_t *len);

#endif