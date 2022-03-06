package volume

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Volume List Provider
func VolumeListProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: volumeDataObj.List,
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
			"volumes": &schema.Schema{
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
						"volume_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_offering_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vm_instance_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_disk_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_offering_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"zone_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_shrink": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_time_stamp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"job_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
