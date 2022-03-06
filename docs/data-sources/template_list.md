# stackbill_template_list

This resource would help us to get either list or a single item of the virtual machine template in the Stackbill.

## Example Usage

```
data "stackbill_sshkey_list" "data_name" {
   uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
   zone_uuid  = "728e9c82-506d-4afb-a26f-3f74688e0740"
}

```

## Argument Reference

The following arguments are supported:

- `uuid` - (Optional) The UUID of the template
- `zone_uuid` - (Required) The UUID of the zone in which templates are created

The following arguments are exported:

- `uuid` : The UUID of the template
- `name` : The name of the template
- `description` : short description about the template
- `osCategoryName` : OS category of template
- `format` : template format
- `zoneName` : name of the zone in which template is created
- `zoneUuid` : UUID of the zone in which template is created
- `templateCost` : template const
- `isActive` : boolean value which describes whether the template is active
