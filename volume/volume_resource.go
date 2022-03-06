package volume

import (
	"context"
	"encoding/json"
	"log"
	"terraform-provider-stackbill/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Volume Resource Object
func NewVolumeResource() VolumeResource {
	return &volumeResource{}
}

//Volume Interface
type VolumeResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Volume Request Object
type volumeResource struct {
}

// Volume Action
// TODO - Documentation
func (vr *volumeResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	log.Println("Volume creation initiated...!")
	volumeRequst := volumeUtilObj.GetVolumeCreateRequest(d)
	response, err := volumeApiObj.CreateVolume(volumeRequst, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	var resJson map[string]interface{}
	if err := json.Unmarshal([]byte(response), &resJson); err != nil {
		return diag.FromErr(err)
	}
	volumeRes := resJson["listVolumeResponse"].([]interface{})
	if len(volumeRes) > 0 {
		volumeObj := volumeRes[0].(map[string]interface{})
		jobId := volumeObj["jobId"].(string)
		// Update the state
		d.SetId(jobId)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Volume creation completed...!")
	return diags
}

// Volume Action
// TODO - Documentation
func (vr *volumeResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Volume Action
// TODO - Documentation
func (vr *volumeResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Volume Action
// TODO - Documentation
func (vr *volumeResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	log.Println("Volume delete initiated...!")
	response, err := volumeApiObj.DeleteVolume(d.Id(), meta)
	if err != nil {
		return diag.FromErr(err)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Volume deleted successfully...!")
	// Reset the terraform state
	d.SetId("")
	return diags
}
