#ifndef __MESSAGE_H
#define __MESSAGE_H

#include <sgx_report.h>
#include <uuid/uuid.h>

typedef struct req_data {
	unsigned char mac[MAC_SIZE];
	unsigned char nonce[16];
	uuid_t cliend_id;
} req_data_t;

typedef struct ra_req_data {
	unsigned char nonce[16];
	uuid_t cliend_id;
} ra_req_data_t;

typedef struct msg_complete_struct {
	uuid_t image_id;
	uuid_t cliend_id;
} msg_cmpt_t;