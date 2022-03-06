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

// Vpc Network Offering Object
func NewVpcNetworkOfferingData() VpcNetworkOfferingData {
	return &vpcnetworkOfferingData{}
}

// Vpc Network Offering interface
type VpcNetworkOfferingData interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Vpc Network Offering Object
type vpcnetworkOfferingData struct {
}

// List Vpcnetworks
// TODO - Documentation
func (vno *vpcnetworkOfferingData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	uuid := d.Get("uuid").(string)
	log.Println("Vpc network offering list initiated...!")
	response, err := networkApiObj.ListVpcNetworkOfferings(uuid, meta)
	log.Println(response)
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
	datas := jsonRes["listNetworkOfferingResponse"].([]interface{})
	list := make([]map[string]interface{}, 0)
	for _, value := range datas {
		v := value.(map[string]interface{})
		i := make(map[string]interface{})
		for k := range v {
			snakeKey := utils.CamelCaseStringToSnakeCase(k)
			i[snakeKey] = v[k]
		}
		list = append(list, i)
	}
	if err := d.Set("vpcnetworkofferings", list); err != nil {
		return diag.FromErr(err)
	}
	log.Println("List vpc offering networks successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
