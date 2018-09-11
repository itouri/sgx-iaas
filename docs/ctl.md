# ctl

iaas

## nova
vm create --image-id [:image_id]
vm delete --image-id [:image_id]
// vm stop --image-id [:image_id]
// vm start --image-id [:image_id]
vm list

## glance
image register [:file_path]
image delete [:image_id]
image crypto [:file_path] [:out_file_path] (interact with RAServer)
image list

## heat
templete register [:file_path]
templete list