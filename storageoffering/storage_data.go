package storageoffering

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

// Storage Offering Object
func NewStorageOfferingData() StorageOfferingData {
	return &storageOfferingData{}
}

// Storage Offering Interface
type StorageOfferingData interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Storage offering Object
type storageOfferingData struct {
}

// Storage Offering List
// TODO - Documentation
func (so *storageOfferingData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	zoneUuid := d.Get("zone_uuid").(string)
	uuid := d.Get("uuid").(string)
	log.Println("Storage Offering list initiated...!")
	response, err := storageOfferingApiObj.ListStorageOfferings(zoneUuid, uuid, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	jsonRes := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(response), &jsonRes); err != nil {
		return diag.FromErr(err)
	}
	count := jsonRes["count"].(float64)
	if err := d.Set("length", count); err != nil {
		return diag.FromErr(err)
	}
	datas := jsonRes["listStorageOfferingResponse"].([]interface{})
	output := make([]map[string]interface{}, 0)
	for _, value := range datas {
		v := value.(map[string]interface{})
		i := make(map[string]interface{})
		for k := range v {
			snakeKey := utils.CamelCaseStringToSnakeCase(k)
			i[snakeKey] = v[k]
		}
		output = append(output, i)
	}
	if err := d.Set("storageofferings", output); err != nil {
		return diag.FromErr(err)
	}
	log.Println("List storage offerings successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
