package instance

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceActionsProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Action Api's",
		CreateContext: instanceActionsResourceObj.Create,
		ReadContext:   instanceActionsResourceObj.Read,
		UpdateContext: instanceActionsResourceObj.Update,
		DeleteContext: instanceActionsResourceObj.Delete,
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
					if v != START && v != STOP {
						errs = append(errs, fmt.Errorf("The action must be %s or %s", START, STOP))
					}
					return
				},
			},
		},
	}
}
