package instance

import (
	"context"
	"encoding/json"
	"log"
	"terraform-provider-stackbill/utils"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Instance Object
func NewInstance() InstanceResource {
	return &instanceResource{}
}

//Instance Interfae
type InstanceResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Instance Request Object
type instanceResource struct {
}

// Create Instance
// TODO - Documentation
func (vs *instanceResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Set all the values
	log.Println("Creating instance request object...!")
	instanceRequest := instanceUtilsObj.GetCreateInstanceRequest(d)
	log.Println("Created instance request object...!")
	log.Println("Instance create request initiated...!")
	response, err := instanceApiObj.CreateInstance(instanceRequest, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	var resJson map[string]interface{}
	if err := json.Unmarshal([]byte(response), &resJson); err != nil {
		return diag.FromErr(err)
	}
	instanceRes := resJson["listInstanceResponse"].([]interface{})
	if len(instanceRes) > 0 {
		instanceObj := instanceRes[0].(map[string]interface{})
		uuid := instanceObj["uuid"].(string)
		// Update the state
		d.SetId(uuid)
		// Check the status
		time.Sleep(20 * time.Second)
		for {
			status, err := instanceApiObj.GetInstanceStatus(uuid, meta)
			if err != nil {
				log.Println(err.Error())
				break
			}
			if status == "RUNNING" {
				break
			}
			time.Sleep(10 * time.Second)
		}
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance creation successful...!")
	return diags
}

// Read Instance
// TODO - Documentation
func (vs *instanceResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Update Instance
// TODO - Documentation
func (vs *instanceResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Check any changes occured in the name
	if d.HasChange("name") {
		log.Println("Update instance name initiated...!")
		updateNameRequest := instanceUtilsObj.GetUpdateInstanceNameRequest(d)
		response, err := instanceApiObj.UpdateInstanceName(updateNameRequest, meta)
		if err != nil {
			return diag.FromErr(err)
		}
		output := utils.FormatJsonString(response)
		log.Println(output)
		log.Println("Update instance name successful...!")
	}
	// Check any changes occured in the compute offering
	if d.HasChange("compute_offering_uuid") {
		log.Println("Instance resize initiated...!")
		resizeRequest := instanceUtilsObj.GetResizeInstanceRequest(d)
		response, err := instanceApiObj.ResizeInstance(resizeRequest, meta)
		if err != nil {
			return diag.FromErr(err)
		}
		output := utils.FormatJsonString(response)
		log.Println(output)
		log.Println("Instance resize successful...!")
	}
	// Check any changes occured in the compute offering
	if d.HasChange("ssh_key_name") {
		log.Println("Instance Reset sshkey initiated...!")
		resetSshkeyRequest := instanceUtilsObj.GetResetSshKeyInstanceRequest(d)
		response, err := instanceApiObj.ResetSshKey(resetSshkeyRequest, meta)
		if err != nil {
			return diag.FromErr(err)
		}
		output := utils.FormatJsonString(response)
		log.Println(output)
		log.Println("Instance Reset sshkey successful...!")
	}
	// Check any changes occured in the compute offering
	if d.HasChange("storage_offering_uuid") {
		log.Println("Changes ----- Changes " + d.Get("storage_offering_uuid").(string))
	}
	return diags
}

// Delete Instance
// TODO - Documentation
func (vs *instanceResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	log.Println("Instance delete request initiated...!")
	uuid := d.Id()
	response, err := instanceApiObj.DeleteInstance(uuid, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance resize successful...!")
	d.SetId("")
	return diags
}
