# stackbill_network_offering_list

This resource would help us to get either list or a single item of the network offering in the Stackbill.

## Example Usage

```
data "stackbill_network_offering_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
   zone_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the network offering
- `zone_uuid` - (Required) The UUID of the zone in which network offerings are created

The following arguments are exported:

- `uuid` : The UUID of the network offering
- `name` : The name of the network offering
- `displayText` : short description about the network offering
- `offerName` : offer name of the network
- `networkTrialOffering` : network trail offering
- `availability` : describes the availability of the network offering
- `guestIpType` : guest IP type of the network offering
- `isActive` : boolean value which describes whether the network offering is active
