# stackbill_instance_resize

This resource would help us to resize the Virtual Machine.

## Example Usage

```
resource "stackbill_instance_resize" "resource_name" {
  uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
  compute_offering_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  cpu_core = "2"
  memory = "1"
}
```

-> The instance should be in a stopped state to use this resource.

## Argument Reference

The following arguments are supported:

- `uuid` - (Required) The UUID of the compute instance in which the action is to be performed.
- `compute_offering_uuid` - (Required) The UUID of the compute offering for the virtual machine.
- `cpu_core` - (Required) no. of CPU core allocated to the virtual machine'
- `memory` - (Required) size of the memory allocated to the virtual machine'
