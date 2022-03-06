# stackbill_instance_list

This resource would help us to get either list or a single item of the instance in the Stackbill.

## Example Usage

```
data "stackbill_instance_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
   zone_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the instance.
- `zone_uuid` - (Required) The UUID of the zone in which instances are created

The following arguments are exported:

- `uuid` : The UUID of the instance
- `name` : The name of the instance
- `displayName` : display name of the instance
- `description` : brief description about the instance
- `diskSize` : disk size of the instance in KB
- `state` : instance current state
- `instanceOwnerName` : Name of the instance's owner
- `zoneUuid` : The UUID of the zone to which the instance is created.
- `networkName` : name of the network attached to the instance
- `templateName` : name of the template which is used to create the instance.
- `cpuCore` : no. of CPU cores in the instance
- `memory` : the size of the memory in MB
- `volumeSize` : volume size of the instance in MB
- `status` : working status of the instance
- `instancePrivateIp` : private IP of the instance
- `networkUuid` : network UUID of the instance
- `storageOfferingUuid` : storage offering UUID of the instance
- `computeOfferingUuid` : compute offering UUID of the instance
- `templateUuid` : template UUID of the instance
- `sshUuid` : UUID of the ssh key used in the instance
- `rootDiskSize` : root disk size of the instance
- `isActive` : boolean value which describes whether the instance is active
