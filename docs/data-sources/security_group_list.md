# stackbill_security_group_list

This resource would help us to get either list or single item of security group in the stackbill.

## Example Usage

```
data "stackbill_security_group_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the security group list

The following arguments are exported:

- `uuid` : The UUID of the security group list
- `name` : The name of the security group list
- `isActive` : boolean value which describes whether the security group is active
- `description` : The description about the security group list
- `status` : The value must be either "ENABLED" or "DISABLED"
- `securityGroupFirewallRule` : Firewall rules of the security group
- `securityGroupPortForwarding` : an object which describes port Forwarding configurations for the security group
- `securityGroupEgressRule` : an object which describes egress rules configured for the security group
