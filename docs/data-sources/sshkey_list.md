# stackbill_sshkey_list

This resource would help us to get either list or a single item of the ssh key in the Stackbill.

## Example Usage

```
data "stackbill_sshkey_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
   zone_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the ssh key
- `zone_uuid` - (Required) The UUID of the zone in which ssh keys are created

The following arguments are exported:

- `uuid` : The UUID of the ssh key
- `name` : The name of the ssh key
- `domainName` : Domain name of the ssh key
- `status` : The status of the key
- `isActive` : boolean value which describes whether the ssh key is active
