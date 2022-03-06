package securitygroup

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SecurityGroupListProvider() *schema.Resource {
	return &schema.Resource{
		ReadContext: securityGroupDataObj.List,
		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"securitygroups": &schema.Schema{
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
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"security_group_firewall_rule": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"end_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"icmp_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"icmp_code": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_group_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"cidr_list": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_active": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"security_group_port_forwarding": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"is_active": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"private_start_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_end_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"public_end_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"public_start_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_group_uuuid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"security_group_egress_rule": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"end_port": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"icmp_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"icmp_code": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_group_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"cidr_list": &schema.Schema{
										Type:     schema.TypeString,
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
				},
			},
		},
	}
}
