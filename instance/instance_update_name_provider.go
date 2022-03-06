package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceUpdateNameProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Action Api's",
		CreateContext: instanceUpdateNameResourceObj.Create,
		ReadContext:   instanceUpdateNameResourceObj.Read,
		UpdateContext: instanceUpdateNameResourceObj.Update,
		DeleteContext: instanceUpdateNameResourceObj.Delete,
		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
