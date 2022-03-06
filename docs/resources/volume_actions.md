# stackbill_volume_actions

This resource would help us to attach or detach the volume to the Virtual Machine

## Example Usage

```
resource "stackbill_volume_actions" "resource_name" {
  uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  instance_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  action = "Attach"
}
```

## Argument Reference

The following arguments are supported:

- `uuid` - (Required) The UUID of the compute instance in which the action is to be performed.
- `instance_uuid` - (Required) The UUID of the volume that is to be attached or detached.
- `action` - (Required) The action must be Attach or Detach.
