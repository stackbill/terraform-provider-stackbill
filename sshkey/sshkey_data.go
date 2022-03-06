package sshkey

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

// Sshkey Object
func NewSshkeyData() SshkeyData {
	return &sshKey{}
}

// Sshkey interface
type SshkeyData interface {
	List(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Sshkey Object
type sshKey struct {
}

// List networks
// TODO - Documentation
func (sk *sshKey) List(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	log.Println("Sshkey list initiated...!")
	response, err := sshkeyApiObj.ListSshkeys(meta)
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
	datas := jsonRes["listSSHKeyResponse"].([]interface{})
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
	if err := d.Set("sshkeys", list); err != nil {
		return diag.FromErr(err)
	}
	log.Println("List Sshkey successful...!")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
