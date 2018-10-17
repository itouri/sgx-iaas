#ifndef __E_FILEIO_H
#define __E_FILEIO_H

#include <sgx_utils.h>

int send_to_ras(char *src, size_t sz);
int recv_from_ras(char **dest, size_t *sz);
int read_file(unsigned char *dest, char *file, off_t *len);

#endif