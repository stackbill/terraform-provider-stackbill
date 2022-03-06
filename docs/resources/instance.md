# stackbill_instance

Provides a Stackbill instance resources

## Example Usage

```
#creating stackbill_instance
resource "stackbill_instance" "instance_name" {
  compute_offering_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  hypervisor_name       = "XenServer"
  memory                =  "512"
  name                  =  "VM_name"
  network_uuid          = "728e9c82-506d-4afb-a26f-3f74688e0740"
  storage_offering_uuid = "728e9c82-506d-4afb-a26f-3f74688e0740"
  template_uuid         = "728e9c82-506d-4afb-a26f-3f74688e0740"
  zone_uuid             = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `zone_uuid` - (Required) The zone that the machine should be created in.
- `name` - (Required) A unique name for the VM,. Changing this forces a new resource to be created.
- `network_uuid` - (Required) The UUID of the network offering to attach to the instance.
- `template_uuid` - (Required) The template UUID.
- `compute_offering_uuid` - (Required) The UUID of the compute offering for creating the VM resources
- `cpu_core` - (optional) no. of CPU core in the case of custom VM offering.
- `disk_size` - (optional) Specify disk size in GB in the case of for the custom VM offering
- `hypervisor_name` - (optional) The hypervisor name
- `memory` - (optional) The memory size of the VM
- `security_group_name` - (optional) The security group name
- `ssh_key_name` - (optional) The ssh key name
- `storage_offering_uuid` - (optional) The UUID of the storage offering for creating the VM resources
