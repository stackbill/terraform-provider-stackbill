package securitygroup

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"terraform-provider-stackbill/utils"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Securitygroup Object
func NewSecuritygroupData() SecuritygroupData {
	return &securitygroupData{}
}

// Securitygroup Interface
type SecuritygroupData interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Securitygroup data object
type securitygroupData struct {
}

// List Securitygroup
// TODO - Documentation
func (co *securitygroupData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	uuid := d.Get("uuid").(string)
	log.Println("Securitygroup list initiated...!")
	response, err := securityGroupApiObj.ListSecurityGroups(uuid, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	jsonRes := make(map[string]interface{})
	if err := json.Unmarshal([]byte(response), &jsonRes); err != nil {
		return diag.FromErr(err)
	}
	count := jsonRes["count"].(float64)
	if err := d.Set("length", count); err != nil {
		return diag.FromErr(err)
	}
	datas := jsonRes["listSecurityGroupResponse"].([]interface{})
	list := make([]map[string]interface{}, 0)
	for _, value := range datas {
		v := value.(map[string]interface{})
		i := make(map[string]interface{})
		for k := range v {
			snakeKey := utils.CamelCaseStringToSnakeCase(k)
			i[snakeKey] = v[k]
		}
		// Port forwarding list
		sgpr := v["securityGroupPortForwarding"].([]interface{})
		sgpList := make([]map[string]interface{}, 0)
		for _, value := range sgpr {
			v := value.(map[string]interface{})
			i := make(map[string]interface{})
			for k := range v {
				snakeKey := utils.CamelCaseStringToSnakeCase(k)
				i[snakeKey] = v[k]
			}
			sgpList = append(sgpList, i)
		}
		i["security_group_port_forwarding"] = sgpList
		// Firewall rule
		sgfr := v["securityGroupFirewallRule"].([]interface{})
		sgfrList := make([]map[string]interface{}, 0)
		for _, value := range sgfr {
			v := value.(map[string]interface{})
			i := make(map[string]interface{})
			for k := range v {
				snakeKey := utils.CamelCaseStringToSnakeCase(k)
				i[snakeKey] = v[k]
			}
			sgpList = append(sgfrList, i)
		}
		i["security_group_firewall_rule"] = sgpList
		// Egress rule
		sger := v["securityGroupEgressRule"].([]interface{})
		sgerList := make([]map[string]interface{}, 0)
		for _, value := range sger {
			v := value.(map[string]interface{})
			i := make(map[string]interface{})
			for k := range v {
				snakeKey := utils.CamelCaseStringToSnakeCase(k)
				i[snakeKey] = v[k]
			}
			sgerList = append(sgerList, i)
		}
		i["security_group_egress_rule"] = sgerList
		list = append(list, i)
	}
	if err := d.Set("securitygroups", list); err != nil {
		return diag.FromErr(err)
	}
	log.Println("Securitygroup list successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
