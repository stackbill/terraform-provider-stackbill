# stackbill_network_list

This resource would help us to get either list or a single item of the network in the Stackbill.

## Example Usage

```
data "stackbill_network_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
   zone_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the network.
- `zone_uuid` - (Required) The UUID of the zone in which networks are created

The following arguments are exported:

- `uuid` : The UUID of the network offering
- `name` : The name of the network offering
- `domainName` : name of the domain in which the network is created
- `cIDR` : CIDR value of the network
- `gateway` : default gateway of the network
- `networkType` : network type
- `networkOfferingUuid` : The UUID of the network offering in which the network is created
- `zoneUuid` : The UUID of the zone to which the network is created.
- `networkDomain` : name of the network domain
- `status` : status of the network
- `cleanUpNetwork`
- `isActive` : boolean value which describes whether the network offering is active
