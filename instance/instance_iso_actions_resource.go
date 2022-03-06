package instance

import (
	"context"
	"log"
	"terraform-provider-stackbill/utils"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Iso Action Resource Object
func NewInstanceIsoActionsResource() InstanceActionsResource {
	return &instanceIsoActionsResource{}
}

// Iso Action interface
type InstanceIsoActionsResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Instance Request Object
type instanceIsoActionsResource struct {
}

// Create Instance
// TODO - Documentation
func (ac *instanceIsoActionsResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	log.Println("Instance iso " + action + " initiated...!")
	isoActionRequest := instanceUtilsObj.GetInstanceIsoActionRequest(d)
	response, err := instanceApiObj.InstanceIsoActions(isoActionRequest, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	// Wait to start or stop the instance
	time.Sleep(20 * time.Second)
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance iso " + action + " completed...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Read Instance
// TODO - Documentation
func (vs *instanceIsoActionsResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Update Instance
// TODO - Documentation
func (vs *instanceIsoActionsResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	log.Println("Instance iso " + action + " initiated...!")
	isoActionRequest := instanceUtilsObj.GetInstanceIsoActionRequest(d)
	response, err := instanceApiObj.InstanceIsoActions(isoActionRequest, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	// Wait to start or stop the instance
	time.Sleep(20 * time.Second)
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance iso " + action + " completed...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Delete Instance
// TODO - Documentation
func (vs *instanceIsoActionsResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Reset the terraform state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
