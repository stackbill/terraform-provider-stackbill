# stackbill_instance_actions

Provides Stackbill instance action resources

## Example Usage

```
resource "stackbill_instance_actions" "status" {
  uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
  action = "Stop"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Required) The UUID of the compute instance in which the action is to be performed.
- `action` - (Required) action must be Start or Stop.
