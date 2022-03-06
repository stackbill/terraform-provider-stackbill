package computeoffering

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

// Compute Offering Object
func NewComputeOfferingData() ComputeOfferingData {
	return &computeOfferingData{}
}

// Compute Offering Interface
type ComputeOfferingData interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Compute Offering Object
type computeOfferingData struct {
}

// List compute Offerings
// TODO - Documentation
func (co *computeOfferingData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	zoneUuid := d.Get("zone_uuid").(string)
	uuid := d.Get("uuid").(string)
	log.Println("Compute Offering list initiated...!")
	response, err := computeOfferingApiObj.ListComputeOfferings(zoneUuid, uuid, meta)
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
	datas := jsonRes["listComputeOfferingResponse"].([]interface{})
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
	if err := d.Set("computeofferings", list); err != nil {
		return diag.FromErr(err)
	}
	log.Println("List compute offerings successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
