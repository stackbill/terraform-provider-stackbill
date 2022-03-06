# stackbill_volume_list

This resource would help us to get either list or a single item of volumes in the Stackbill.

## Example Usage

```
data "stackbill_volume_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
   zone_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the volume
- `zone_uuid` - (Required) The UUID of the zone in which volumes are created

The following arguments are exported:

- `uuid` : The UUID of the volume
- `name` : The name of the volume
- `volumeType` : Type of the volume
- `storageOfferingName` : The storage offering in which the volume is created
- `storageOfferingUuid` : The storage offering UUID in which the volume is created
- `zoneUuid` : The UUID of the zone in which volumes are created
- `vmInstanceName` : The name of the VM the volume is attached.
- `storageDiskSize` : The size of the volume in kb
- `status` : The status of the volume
- `isActive` : boolean value which describes whether the volume is active
