package network

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NetworkActionProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Network Actions Api's",
		CreateContext: networkActionResourceObj.Create,
		ReadContext:   networkActionResourceObj.Read,
		UpdateContext: networkActionResourceObj.Update,
		DeleteContext: networkActionResourceObj.Delete,
		Schema: map[string]*schema.Schema{
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"virutal_machine_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v != ADD && v != DELETE {
						errs = append(errs, fmt.Errorf("The action must be %s or %s", ADD, DELETE))
					}
					return
				},
			},
		},
	}
}
