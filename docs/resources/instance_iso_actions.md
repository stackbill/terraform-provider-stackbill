# stackbill_instance_iso_actions

Provides Stackbill instance iso actions to attach or detach ISO from the Virtual Machine.

## Example Usage

```
resource "stackbill_instance_actions" "status" {
  iso_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  action = "Detach"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Required) The UUID of the compute instance in which the action is to be performed.
- `iso_uuid` - (Optional) The UUID of the ISO that is to be attached to the Virtual Machine.
- `action` - (Required) action must be Attach or Detach.
