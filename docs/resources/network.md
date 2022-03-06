# stackbill_network

This resource would help us to create the network in the Stackbill.

## Example Usage

```
resource "stackbill_network" "TerraformNetwork" {
  description = "short description about network"
  name = "network_name"
  virtual_machine_uuid = "19c28a5d-0237-450d-b47a-dbddc67aa0df"
  network_offering_uuid = "19c28a5d-0237-450d-b47a-dbddc67aa0df"
  zone_uuid = "19c28a5d-0237-450d-b47a-dbddc67aa0df"
  is_public = "true"
}
```

## Argument Reference

The following arguments are supported:

- `virtual_machine_uuid ` - (Required) The UUID of the compute instance in which the action is to be performed.
- `description ` - (Required) short description about the network
- `name ` - (Required) Name for the new network.
- `zone_uuid` - (Required) The UUID of the zone that the network should be created in.
- `network_offering_uuid` - (Required) The UUID of the network offering in which the network is created.
- `is_public` - (Required) The value should be either "true" or "false"
