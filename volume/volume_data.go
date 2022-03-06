package volume

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

// Volume Object
func NewVolumeData() VolumeData {
	return &volumeData{}
}

// Volume Interface
type VolumeData interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Volume Data
type volumeData struct {
}

// List Zones
// TODO - Documentation
func (v *volumeData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get the Uuid
	uuid := d.Get("uuid").(string)
	zoneUuid := d.Get("zone_uuid").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	log.Println("Volume list initiated...!")
	response, err := volumeApiObj.ListVolumes(zoneUuid, uuid, meta)
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
	datas := jsonRes["listVolumeResponse"].([]interface{})
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
	if err := d.Set("volumes", list); err != nil {
		return diag.FromErr(err)
	}
	log.Println("List volumes successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
