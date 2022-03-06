# stackbill_vpc_network_offering_list

This resource would help us to get either list or a single item of the VPC network offering in the Stackbill.

## Example Usage

```
data "stackbill_vpc_network_offering_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the vpc network offering

The following arguments are exported:

- `uuid` : The UUID of the vpc network offering
- `name` : The name of the vpc network offering
- `displayText` : short description about the vpc network offering
- `offerName` : offer name of the vpc network
- `networkTrialOffering` : network trail offering
- `availability` : describes the availability of the vpc network offering
- `guestIpType` : guest IP type of the vpc network offering
- `isActive` : boolean value which describes whether the vpc network offering is active
