# stackbill_instance_update_name

Provides resources to update the display name of the Virtual Machine.

## Example Usage

```
resource "stackbill_instance_update_name" "resource_name" {
  uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
  name = "New_vm_display_name"
}
```

-> The instance should be in a stopped state to use this resource.

## Argument Reference

The following arguments are supported:

- `uuid` - (Required) The UUID of the compute instance in which the action is to be performed.
- `name` - (Required) new display name for the Virtual Machine.
