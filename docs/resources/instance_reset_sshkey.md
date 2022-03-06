# stackbill_instance_reset_sshkey

This resource would help us to reset the ssh key of the Virtual Machine

## Example Usage

```
resource "stackbill_instance_reset_sshkey" "resource_name" {
  uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  ssh_key_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
}
```

## Argument Reference

The following arguments are supported:

- `uuid` - (Required) The UUID of the compute instance in which the action is to be performed.
- `ssh_key_uuid` - (Required) The UUID of the SSH key that is to be replaced.
