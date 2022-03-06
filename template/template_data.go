package template

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
func NewTemplateData() TemplateData {
	return &templateData{}
}

// Template Interface
type TemplateData interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Template data object
type templateData struct {
}

// List Templates
// TODO - Documentation
func (co *templateData) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	zoneUuid := d.Get("zone_uuid").(string)
	uuid := d.Get("uuid").(string)
	log.Println("Template list initiated...!")
	response, err := templateApiObj.ListTemplates(zoneUuid, uuid, meta)
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
	if err := d.Set("templates", output); err != nil {
		return diag.FromErr(err)
	}
	log.Println("List Templates successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
