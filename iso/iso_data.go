package iso

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

// Template Object
func NewIsoData() IsoData {
	return &isoData{}
}

// Template Interface
type IsoData interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Iso data object
type isoData struct {
}

// List Templates
// TODO - Documentation
func (co *isoData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	zoneUuid := d.Get("zone_uuid").(string)
	uuid := d.Get("uuid").(string)
	log.Println("Iso list initiated...!")
	response, err := isoApiObj.ListIso(zoneUuid, uuid, meta)
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
	datas := jsonRes["listTemplateResponse"].([]interface{})
	list := make([]map[string]interface{}, 0)
	for _, value := range datas {
		v := value.(map[string]interface{})
		i := make(map[string]interface{})
		format := v["format"].(string)
		if format != "ISO" {
			continue
		}
		for k := range v {
			snakeKey := utils.CamelCaseStringToSnakeCase(k)
			i[snakeKey] = v[k]
		}
		list = append(list, i)
	}
	if err := d.Set("iso", list); err != nil {
		return diag.FromErr(err)
	}
	log.Println("Iso list successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
