# stackbill_compute_offering_list

This resource would help us to get either list or a single item of the compute offering in the Stackbill.

## Example Usage

```
data "stackbill_compute_offering_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
   zone_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the compute offering.
- `zone_uuid` - (Required) The UUID of the zone in which compute offerings are created

The following arguments are exported:

- `uuid` : The UUID of the compute offering
- `name` : The name of the compute offering
- `displayText` : the display name of the compute offering
- `numberOfCores` : no. of CPU cores allocated by the compute offering
- `clockSpeed` : clock speed allocated by compute offering
- `memory` : size of the memory allocated by the compute offering
- `storageType` : storage type provided by the offering to the instance created with this compute offering
- `isActive` : boolean value which describes whether the compute offering is active
