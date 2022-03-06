package computeoffering

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ComputeOfferingListProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: computeOfferingDataObj.List,
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
			"computeofferings": &schema.Schema{
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
						"display_text": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"number_of_cores": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"clock_speed": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
