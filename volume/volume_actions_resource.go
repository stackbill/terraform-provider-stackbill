package volume

import (
	"context"
	"encoding/json"
	"log"
	"terraform-provider-stackbill/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Action Resource Object
func NewVolumeActionsResource() VolumeActionsResource {
	return &volumeActionsResource{}
}

//Volume Interfae
type VolumeActionsResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Volume Request Object
type volumeActionsResource struct {
}

// Volume Action
// TODO - Documentation
func (va *volumeActionsResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	log.Println("Volume " + action + " initiated...!")
	actionRequest := volumeUtilObj.GetVolumeActionRequest(d)
	response, err := volumeApiObj.VolumeActions(actionRequest, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	var resJson map[string]interface{}
	if err := json.Unmarshal([]byte(response), &resJson); err != nil {
		return diag.FromErr(err)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance " + action + " completed...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Volume Action
// TODO - Documentation
func (va *volumeActionsResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Volume Action
// TODO - Documentation
func (va *volumeActionsResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	log.Println("Volume " + action + " initiated...!")
	actionRequest := volumeUtilObj.GetVolumeActionRequest(d)
	response, err := volumeApiObj.VolumeActions(actionRequest, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance " + action + " completed...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Volume Action
// TODO - Documentation
func (va *volumeActionsResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Reset the terraform state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
