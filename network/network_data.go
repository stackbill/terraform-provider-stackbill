package network

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

// Network Object
func NewNetworkData() NetworkData {
	return &networkData{}
}

// Network interface
type NetworkData interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Network Object
type networkData struct {
}

// List networks
// TODO - Documentation
func (co *networkData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	zoneUuid := d.Get("zone_uuid").(string)
	uuid := d.Get("uuid").(string)
	log.Println("Network list initiated...!")
	response, err := networkApiObj.ListNetworks(zoneUuid, uuid, meta)
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
	datas := jsonRes["listNetworkResponse"].([]interface{})
	list := make([]map[string]interface{}, 0)
	for _, value := range datas {
		v := value.(map[string]interface{})
		i := make(map[string]interface{})
		for k := range v {
			snakeKey := utils.CamelCaseStringToSnakeCase(k)
			if k == "cIDR" {
				i["cidr"] = v[k]
				continue
			}
			i[snakeKey] = v[k]
		}
		list = append(list, i)
	}
	if err := d.Set("networks", list); err != nil {
		return diag.FromErr(err)
	}
	log.Println("List networks successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
