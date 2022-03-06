package instance

import (
	"context"
	"log"
	"strings"
	"terraform-provider-stackbill/utils"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ACtion Resource Object
func NewInstanceActionsResource() InstanceActionsResource {
	return &instanceActionsResource{}
}

//Instance Interfae
type InstanceActionsResource interface {
	Create(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Read(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Update(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
	Delete(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
}

//Instance Request Object
type instanceActionsResource struct {
}

// Create Instance
// TODO - Documentation
func (vs *instanceActionsResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	log.Println("Instance " + action + " initiated...!")
	actionRequest := instanceUtilsObj.GetInstanceActionRequest(d)
	response, err := instanceApiObj.InstanceActions(actionRequest, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	// Wait to start or stop the instance
	time.Sleep(20 * time.Second)
	for {
		status, err := instanceApiObj.GetInstanceStatus(uuid, meta)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if action == START && status == "RUNNING" {
			response = strings.Replace(response, "STOPPED", "RUNNING", -1)
			break
		}
		if action == STOP && status == "STOPPED" {
			response = strings.Replace(response, "RUNNING", "STOPPED", -1)
			break
		}
		time.Sleep(10 * time.Second)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance " + action + " completed...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Read Instance
// TODO - Documentation
func (vs *instanceActionsResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

// Update Instance
// TODO - Documentation
func (vs *instanceActionsResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Get name and object
	uuid := d.Get("uuid").(string)
	action := d.Get("action").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// Action
	log.Println("Instance " + action + " initiated...!")
	actionRequest := instanceUtilsObj.GetInstanceActionRequest(d)
	response, err := instanceApiObj.InstanceActions(actionRequest, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	time.Sleep(20 * time.Second)
	for {
		status, err := instanceApiObj.GetInstanceStatus(uuid, meta)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if action == START && status == "RUNNING" {
			response = strings.Replace(response, "STOPPED", "RUNNING", -1)
			break
		}
		if action == STOP && status == "STOPPED" {
			response = strings.Replace(response, "RUNNING", "STOPPED", -1)
			break
		}
		time.Sleep(10 * time.Second)
	}
	output := utils.FormatJsonString(response)
	log.Println(output)
	log.Println("Instance " + action + " completed...!")
	// Update the state
	d.SetId(uuid)
	return diags
}

// Delete Instance
// TODO - Documentation
func (vs *instanceActionsResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// log.Println("Instance action detach initiated...!")
	// actionRequest := instanceUtilsObj.GetInstanceActionRequest(d)
	// actionRequest.Action = STOP
	// response, err := instanceApiObj.InstanceActions(actionRequest, meta)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }
	// output := utils.FormatJsonString(response)
	// log.Println(output)
	// log.Println("Instance action detach completed...!")
	// // Update the state
	d.SetId("")
	return diags
}
