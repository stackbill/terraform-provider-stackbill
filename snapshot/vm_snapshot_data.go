package snapshot

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

// Vm Snapshot Data
func NewVmSnapshotData() VmSnapshotData {
	return &vmSnapshotData{}
}

// Vm Snapshot Data Interface
type VmSnapshotData interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Storage offering Object
type vmSnapshotData struct {
}

// Storage Offering List
// TODO - Documentation
func (vs *vmSnapshotData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	zoneUuid := d.Get("zone_uuid").(string)
	uuid := d.Get("uuid").(string)
	log.Println("Vm Snapshot list initiated...!")
	response, err := vmSnapshotApiObj.ListVmSnapshots(zoneUuid, uuid, meta)
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
	datas := jsonRes["listVmSnapshotResponse"].([]interface{})
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
	if err := d.Set("vmsnapshots", list); err != nil {
		return diag.FromErr(err)
	}
	log.Println("List Vm Snapshot successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
