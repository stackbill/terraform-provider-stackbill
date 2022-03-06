# stackbill_iso_list

This resource would help us to get either list or a single item of the ISO in the Stackbill.

## Example Usage

```
data "stackbill_iso_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
   zone_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the iso
- `zone_uuid` - (Required) The UUID of the zone in which iso is created
