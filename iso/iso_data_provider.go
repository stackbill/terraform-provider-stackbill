package iso

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func IsoListProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: isoDataObj.List,
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
			"iso": &schema.Schema{
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
						"os_category_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"zone_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"zone_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_cost": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
