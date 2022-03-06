# stackbill_instance_snapshot

This resource would help us to snapshot the Virtual Machine

## Example Usage

```
resource "stackbill_instance_snapshot" "resource_name" {
  name = "snapshot name"
  description = "snapshot description"
  virtual_machine_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  zone_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  snapshot_memory = "0"
}
```

## Argument Reference

The following arguments are supported:

- `virtual_machine_uuid` - (Required) The UUID of the compute instance in which the action is to be performed.
- `name ` - (Required) Name for the snapshot.
- `description ` - (Required) short description about the snapshot
- `zone_uuid` - (Required) The UUID of the zone that the snapshot should be created in.
- `memory` - (Required) size of the memory allocated to the snapshot'
