package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Api's",
		CreateContext: instanceObj.Create,
		ReadContext:   instanceObj.Read,
		UpdateContext: instanceObj.Update,
		DeleteContext: instanceObj.Delete,
		Schema: map[string]*schema.Schema{
			"compute_offering_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_core": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disk_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hypervisor_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"security_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_key_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_offering_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"template_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"zone_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
