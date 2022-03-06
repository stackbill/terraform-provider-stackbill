# stackbill_storage_offering_list

This resource would help us to get either list or a single item of the storage offering in the Stackbill.

## Example Usage

```
data "stackbill_storage_offering_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
   zone_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the storage offering
- `zone_uuid` - (Required) The UUID of the zone in which storage offerings are created

The following arguments are exported:

- `uuid` : The UUID of the storage offering
- `name` : The name of the storage offering
- `description` : short description about the storage offering
- `storageType` : storage type of the storage offering
- `diskSize` : size of the storage in GB
- `isCustomDisk` : boolean value which describes whether the storage is customizable
- `isActive` : boolean value which describes whether the storage offering is active
