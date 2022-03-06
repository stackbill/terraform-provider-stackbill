# stackbill_volume

This resource would help us to create a new volume in the Stackbill.

## Example Usage

```
resource "stackbill_volume" "newvolume" {
  disk_size = 10
  name = "volume_name"
  storage_offering_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  zone_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `name ` - (Required) Name for the new network.
- `zone_uuid` - (Required) The UUID of the zone that the volume should be created in.
- `storage_offering_uuid` - (Required) The UUID of the storage offering in which the volume should be created in.
- `disk_size ` - (Required) The size of the disk in GB.
