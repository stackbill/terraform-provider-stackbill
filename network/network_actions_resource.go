package network

import (
	"context"
	"encoding/json"
	"log"
	"terraform-provider-stackbill/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Netwok Resource Object
func NewNetworkActionResource() NetworkActionResource {
	return &networkActionResource{}
}

// NetworkAction Interface
type NetworkActionResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// NetworkAction Object
type networkActionResource struct {
}

// Snapshot Instance
// TODO - Documentation
func (nr *networkActionResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("virutal_machine_uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	log.Println("Network " + action + " initiated...!")
	actionRequest := networkUtilObj.GetNetworkActionRequest(d)
	response, err := networkApiObj.NetworkActions(action, actionRequest, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	var resJson map[string]interface{}
	if err := json.Unmarshal([]byte(response), &resJson); err != nil {
		return diag.FromErr(err)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Network " + action + " completed...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Resize Instance
// TODO - Documentation
func (nr *networkActionResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Snapshot Instance
// TODO - Documentation
func (vs *networkActionResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("virutal_machine_uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	log.Println("Network " + action + " initiated...!")
	actionRequest := networkUtilObj.GetNetworkActionRequest(d)
	response, err := networkApiObj.NetworkActions(action, actionRequest, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	var resJson map[string]interface{}
	if err := json.Unmarshal([]byte(response), &resJson); err != nil {
		return diag.FromErr(err)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Network " + action + " completed...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Delete Snapshot
// TODO - Documentation
func (nr *networkActionResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
