package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InstanceResetSshKeyProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance Reset Ssh Key Api's",
		CreateContext: instanceResetSshkeyResourceObj.Create,
		ReadContext:   instanceResetSshkeyResourceObj.Read,
		UpdateContext: instanceResetSshkeyResourceObj.Update,
		DeleteContext: instanceResetSshkeyResourceObj.Delete,
		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ssh_key_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
