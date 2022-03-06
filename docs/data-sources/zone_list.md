# stackbill_zone_list

This resource would help us to get either list or a single item of the zone in the Stackbill.

## Example Usage

```
data "stackbill_zone_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the zone

The following arguments are exported:

- `uuid` : The UUID of the zone
- `name` : The name of the zone
- `isActive` : boolean value which describes whether the zone is active
