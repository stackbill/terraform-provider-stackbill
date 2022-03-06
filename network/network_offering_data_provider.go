package network

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NetworkOfferingListProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: networkOfferingDataObj.List,
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
			"networkofferings": &schema.Schema{
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
						"offer_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"network_trial_offering": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"availability": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"guest_ip_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
