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
func NewNetworkResource() NetworkResource {
	return &networkResource{}
}

// Network Interface
type NetworkResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Network Object
type networkResource struct {
}

// Snapshot Instance
// TODO - Documentation
func (nr *networkResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// Action
	log.Println("Network creation initiated...!")
	networkCreateRequest := networkUtilObj.GetCreateNetworkRequest(d)
	optionalRequest := networkUtilObj.GetCreateNetworkOptionalRequest(d)
	response, err := networkApiObj.CreateNetwork(networkCreateRequest, optionalRequest, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	var resJson map[string]interface{}
	if err := json.Unmarshal([]byte(response), &resJson); err != nil {
		return diag.FromErr(err)
	}
	networkRes := resJson["listNetworkResponse"].([]interface{})
	if len(networkRes) > 0 {
		networObj := networkRes[0].(map[string]interface{})
		uuid := networObj["uuid"].(string)
		// Update the state
		d.SetId(uuid)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Network creation successful...!")
	return diags
}

// Resize Instance
// TODO - Documentation
func (nr *networkResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Snapshot Instance
// TODO - Documentation
func (vs *networkResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Delete Snapshot
// TODO - Documentation
func (nr *networkResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	log.Println("Network delete initiated...!")
	response, err := networkApiObj.DeleteNetwork(d.Id(), meta)
	if err != nil {
		return diag.FromErr(err)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Network deleted successfully...!")
	// Reset the terraform state
	d.SetId("")
	return diags
}
