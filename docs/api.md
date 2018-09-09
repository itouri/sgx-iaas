# api

## ceilometer
GET /alarm
GET /alarm/:alarm_id
POST /alarm
DELETE /alarm/:alarm_id

## glance
GET /images/:image_id
GET /images
POST /images
DELETE /images/:image_id

## heat
GET /stacks/:stack_id
GET /stacks
POST /stacks
DELETE /stacks/:stack_id

## keystone(endpoint)
GET /services/resolve/:service_type

GET /services/:service_id
GET /services
POST /services
DELETE /services/:service_id

// EnumServiceTypeでサービスタイプを指定してIPとポートを返してもらう
GET /services/resolve/:service_type

## neutron
多くね

// floatingips
GET /floatingips
GET /floatingips/:floatingip_id
POST /floatingips
PUT /floatingips/:floatingip_id
DELETE /floatingips/:floatingip_id

// network
GET /networks
GET /networks/:network_id
POST /networks
PUT /networks/:network_id
DELETE /networks/:network_id

// router
GET /routers
GET /routers/:router_id
POST /routers
PUT /routers/:router_id
DELETE /routers/:router_id

// stacks
GET /stacks
GET /stacks/:stack_id
POST /stacks

## nova
POST /vm/:image_id/create

GET /vm/status/:image_id
GET /vm/status

## compute
POST /vm/:image_id/create
