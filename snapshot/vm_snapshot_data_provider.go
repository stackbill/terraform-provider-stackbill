package snapshot

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func VmSnapshotListProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: vmSnapshotDataObj.List,
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
			"vmsnapshots": &schema.Schema{
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
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"domain_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_current": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_time_stamp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"zone_uuid": &schema.Schema{
							Type:     schema.TypeString,
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
