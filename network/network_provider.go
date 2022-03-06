package network

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NetworkProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Snapshot Api's",
		CreateContext: networkResourceObj.Create,
		ReadContext:   networkResourceObj.Read,
		UpdateContext: networkResourceObj.Update,
		DeleteContext: networkResourceObj.Delete,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_machine_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_offering_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"zone_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"is_public": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"security_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
