package snapshot

import (
	"context"
	"encoding/json"
	"log"
	"terraform-provider-stackbill/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Instance Snapshot Object
func NewVmSnapshotResource() VmSnapshotResource {
	return &vmSnapshotResource{}
}

// Snapshot Interfae
type VmSnapshotResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

// Snapshot Object
type vmSnapshotResource struct {
}

// Snapshot Instance
// TODO - Documentation
func (vs *vmSnapshotResource) Snapshot(ctx context.Context, d *schema.ResourceData, meta interface{}) (string, error) {
	// Get the Requst Object
	vmSnapshotRequest := vmSnapshotUtilsObj.GetVmSnapshotRequest(d)
	response, err := vmSnapshotApiObj.SnapshotVm(vmSnapshotRequest, meta)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Snapshot Instance
// TODO - Documentation
func (vs *vmSnapshotResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// Action
	log.Println("Instance snapshot initiated...!")
	response, err := vs.Snapshot(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	var resJson map[string]interface{}
	if err := json.Unmarshal([]byte(response), &resJson); err != nil {
		return diag.FromErr(err)
	}
	instanceRes := resJson["listVmSnapshotResponse"].([]interface{})
	if len(instanceRes) > 0 {
		instanceObj := instanceRes[0].(map[string]interface{})
		jobId := instanceObj["jobId"].(string)
		// Update the state
		d.SetId(jobId)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance snapshot successful...!")
	return diags
}

// Resize Instance
// TODO - Documentation
func (vs *vmSnapshotResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Snapshot Instance
// TODO - Documentation
func (vs *vmSnapshotResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Delete Snapshot
// TODO - Documentation
func (vs *vmSnapshotResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	log.Println("Instance snapshot delete initiated...!")
	response, err := vmSnapshotApiObj.DeleteVmSnapshot(d.Id(), meta)
	if err != nil {
		return diag.FromErr(err)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance snapshot deleted successfully...!")
	// Reset the terraform state
	d.SetId("")
	return diags
}
