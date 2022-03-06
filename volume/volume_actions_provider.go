package volume

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func VolumeActionsProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Volume Action Api's",
		CreateContext: volumeActionsResourceObj.Create,
		ReadContext:   volumeActionsResourceObj.Read,
		UpdateContext: volumeActionsResourceObj.Update,
		DeleteContext: volumeActionsResourceObj.Delete,
		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v != ATTACH && v != DETACH {
						errs = append(errs, fmt.Errorf("The action must be %s or %s", ATTACH, DETACH))
					}
					return
				},
			},
			"instance_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
