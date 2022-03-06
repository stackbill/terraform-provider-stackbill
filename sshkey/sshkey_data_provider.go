package sshkey

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SshkeyListProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: sshkeyDataObj.List,
		Schema: map[string]*schema.Schema{
			"length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sshkeys": &schema.Schema{
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
						"is_active": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"domain_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
