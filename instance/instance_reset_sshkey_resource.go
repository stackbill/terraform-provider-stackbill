package instance

import (
	"context"
	"log"
	"terraform-provider-stackbill/utils"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Instance Resest ssh key Object
func NewInstanceResetSshKeyResource() InstanceResetSshkeyResource {
	return &instanceResetSshkeyResource{}
}

//Instance Interfae
type InstanceResetSshkeyResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Instance Request Object
type instanceResetSshkeyResource struct {
}

// Create Instance
// TODO - Documentation
func (irs *instanceResetSshkeyResource) ResetSshkey(ctx context.Context, d *schema.ResourceData, meta interface{}) (string, error) {
	// Get the Requst Object
	resetSshkeyRequest := instanceUtilsObj.GetResetSshKeyInstanceRequest(d)
	response, err := instanceApiObj.ResetSshKey(resetSshkeyRequest, meta)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Create Instance
// TODO - Documentation
func (irs *instanceResetSshkeyResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	log.Println("Instance Reset sshkey initiated...!")
	response, err := irs.ResetSshkey(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	// Wait to start or stop the instance
	time.Sleep(20 * time.Second)
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance Reset sshkey successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Read Instance
// TODO - Documentation
func (irs *instanceResetSshkeyResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Update Instance
// TODO - Documentation
func (irs *instanceResetSshkeyResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	var diags diag.Diagnostics
	// Action
	log.Println("Instance Reset sshkey initiated...!")
	response, err := irs.ResetSshkey(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	// Wait to start or stop the instance
	time.Sleep(20 * time.Second)
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance Reset sshkey successful...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Delete Instance
// TODO - Documentation
func (irs *instanceResetSshkeyResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Reset the terraform state
	d.SetId("")
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
