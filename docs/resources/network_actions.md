# stackbill_network_actions

This resource would help us to Add or Delete the network to the Virtual Machine

## Example Usage

```
resource "stackbill_network_actions" "attach_network" {
  network_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  virutal_machine_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  action = "Add"
}
```

## Argument Reference

The following arguments are supported:

- `virutal_machine_uuid` - (Required) The UUID of the compute instance in which the action is to be performed.
- `network_uuid` - (Required) The UUID of the network that is to be added or deleted.
- `action` - (Required) action must be Add or Delete
