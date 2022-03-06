package storageoffering

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func StorageOfferingListProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: storageOfferingDataObj.List,
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
			"storageofferings": &schema.Schema{
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
						"storage_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"disk_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_custom_disk": &schema.Schema{
							Type:     schema.TypeBool,
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
