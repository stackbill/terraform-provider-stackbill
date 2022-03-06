package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceListProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: instanceDataObj.List,
		Schema: map[string]*schema.Schema{
			"zone_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"instances": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"disk_size": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"instance_owner_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cpu_core": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_private_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_offering_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"compute_offering_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"zone_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssh_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"root_disk_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
