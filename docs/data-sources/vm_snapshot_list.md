# stackbill_snapshot_list

This resource would help us to get either list or a single item of the snapshot in the Stackbill.

## Example Usage

```
data "stackbill_vm_snapshot_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
   zone_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the snapshot
- `zone_uuid` - (Required) The UUID of the zone in which snapshot are created

The following arguments are exported:

- `uuid` : The UUID of the snapshot
- `name` : The name of the snapshot
- `snapshotType` : Type of the snapshot
- `domainName` : Domain name of the snapshot
- `snapshotTime` : The time when the snapshot is created.
- `volumeUuid` : The volume UUID of the created snapshot.
- `zoneUuid` : The UUID of the zone in which snapshot is created
- `status` : The status of the snapshot
- `isActive` : boolean value which describes whether the snapshot is active
