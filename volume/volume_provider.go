package volume

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func VolumeProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Volume create Api's",
		CreateContext: volumeResourceObj.Create,
		ReadContext:   volumeResourceObj.Read,
		UpdateContext: volumeResourceObj.Update,
		DeleteContext: volumeResourceObj.Delete,
		Schema: map[string]*schema.Schema{
			"disk_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"storage_offering_uuid": &schema.Schema{
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
