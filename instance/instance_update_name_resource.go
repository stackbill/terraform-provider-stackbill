package instance

import (
	"context"
	"log"
	"terraform-provider-stackbill/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Instance Update name Object
func NewInstanceUpdaeNameResource() InstanceUpdateNameResource {
	return &instanceUpdateNameResource{}
}

//Instance Interfae
type InstanceUpdateNameResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Instance Request Object
type instanceUpdateNameResource struct {
}

// Create Instance
// TODO - Documentation
func (iu *instanceUpdateNameResource) UpdateName(ctx context.Context, d *schema.ResourceData, meta interface{}) (string, error) {
	// Get the Requst Object
	updateNameRequest := instanceUtilsObj.GetUpdateInstanceNameRequest(d)
	response, err := instanceApiObj.UpdateInstanceName(updateNameRequest, meta)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Create Instance
// TODO - Documentation
func (iu *instanceUpdateNameResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	log.Println("Update instance name initiated...!")
	response, err := iu.UpdateName(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Update instance name successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Read Instance
// TODO - Documentation
func (iu *instanceUpdateNameResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Update Instance
// TODO - Documentation
func (iu *instanceUpdateNameResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	log.Println("Update instance name initiated...!")
	response, err := iu.UpdateName(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Update instance name successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Delete Instance
// TODO - Documentation
func (iu *instanceUpdateNameResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Reset the terraform state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
