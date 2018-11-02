# api

## glance
GET /images/:image_id 
(GET /images)
POST /images/:image_id
(DELETE /images/:image_id)

<!-- ## nova
POST /vm/create/image -->

## compute
POST /vm/create/:image_id
(POST /vm/stop/:image_id)
(DELETE /vm/delete/:image_id)

## ra
GET /ra/client_id <!-- OK -->
(DELETE /ra/client_id)
POST /ra/images/:client_id 
GET /ra/crypto_key <!-- OK -->
<!-- GET /ra/verify_req/:image_id/:client_id // localで呼べばいい -->