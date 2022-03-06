package instance

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceIsoActionsProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Action Api's",
		CreateContext: instanceIsoActionsResourceObj.Create,
		ReadContext:   instanceIsoActionsResourceObj.Read,
		UpdateContext: instanceIsoActionsResourceObj.Update,
		DeleteContext: instanceIsoActionsResourceObj.Delete,
		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"iso_uuid": &schema.Schema{
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
		},
	}
}
