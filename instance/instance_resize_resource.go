package instance

import (
	"context"
	"log"
	"terraform-provider-stackbill/utils"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Instance Resize Object
func NewInstanceResizeResource() InstanceResizeResource {
	return &instanceResizeResource{}
}

// Resize Interfae
type InstanceResizeResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Resize Object
type instanceResizeResource struct {
}

// Resize Instance
// TODO - Documentation
func (ir *instanceResizeResource) Resize(ctx context.Context, d *schema.ResourceData, meta interface{}) (string, error) {
	// Get the Requst Object
	resizeRequest := instanceUtilsObj.GetResizeInstanceRequest(d)
	response, err := instanceApiObj.ResizeInstance(resizeRequest, meta)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Resize Instance
// TODO - Documentation
func (ir *instanceResizeResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	log.Println("Instance resize initiated...!")
	response, err := ir.Resize(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	// Wait to start or stop the instance
	time.Sleep(2 * time.Second)
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance resize successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Resize Instance
// TODO - Documentation
func (ir *instanceResizeResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Resize Instance
// TODO - Documentation
func (ir *instanceResizeResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	log.Println("Instance resize initiated...!")
	response, err := ir.Resize(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	// Wait to start or stop the instance
	time.Sleep(20 * time.Second)
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance resize successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Delete Instance
// TODO - Documentation
func (ir *instanceResizeResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Reset the terraform state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
