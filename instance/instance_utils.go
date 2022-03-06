package instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Instance util Object
func NewInstanceUtils() InstanceUtils {
	return &instanceUtils{}
}

//Instance Interfae
type InstanceUtils interface {
	GetCreateInstanceRequest(d *schema.ResourceData) map[string]interface{}
	GetUpdateInstanceNameRequest(d *schema.ResourceData) map[string]interface{}
	GetResetSshKeyInstanceRequest(d *schema.ResourceData) map[string]interface{}
	GetResizeInstanceRequest(d *schema.ResourceData) map[string]interface{}
	GetInstanceActionRequest(d *schema.ResourceData) map[string]interface{}
	GetInstanceIsoActionRequest(d *schema.ResourceData) map[string]interface{}
}

//Instance utils Object
type instanceUtils struct {
}

// Create Instance
// TODO - Documentation
func (vs *instanceUtils) GetCreateInstanceRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["computeOfferingUuid"] = d.Get("compute_offering_uuid").(string)
	if d.Get("cpu_core").(string) != "" {
		request["cpuCore"] = d.Get("cpu_core").(string)
	}
	if d.Get("disk_size").(int) >= 0 {
		request["diskSize"] = d.Get("disk_size").(int)
	}
	if d.Get("hypervisor_name").(string) != "" {
		request["hypervisorName"] = d.Get("hypervisor_name").(string)
	}
	if d.Get("memory").(string) != "" {
		request["memory"] = d.Get("memory").(string)
	}
	request["name"] = d.Get("name").(string)
	request["networkUuid"] = d.Get("network_uuid").(string)
	if d.Get("security_group_name").(string) != "" {
		request["securitygroupName"] = d.Get("security_group_name").(string)
	}
	if d.Get("ssh_key_name").(string) != "" {
		request["sshKeyName"] = d.Get("ssh_key_name").(string)
	}
	if d.Get("storage_offering_uuid").(string) != "" {
		request["storageOfferingUuid"] = d.Get("storage_offering_uuid").(string)
	}
	request["templateUuid"] = d.Get("template_uuid").(string)
	request["zoneUuid"] = d.Get("zone_uuid").(string)
	return request
}

// Update name request
// TODO - Documentation
func (vs *instanceUtils) GetUpdateInstanceNameRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["displayName"] = d.Get("name").(string)
	uuid := d.Id()
	if d.Id() == "" {
		uuid = d.Get("uuid").(string)
	}
	request["uuid"] = uuid
	return request
}

// Update name request
// TODO - Documentation
func (vs *instanceUtils) GetResetSshKeyInstanceRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["sshKeyId"] = d.Get("ssh_key_uuid").(string)
	uuid := d.Id()
	if d.Id() == "" {
		uuid = d.Get("uuid").(string)
	}
	request["uuid"] = uuid
	return request
}

// Resize request
// TODO - Documentation
func (vs *instanceUtils) GetResizeInstanceRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["offeringUuid"] = d.Get("compute_offering_uuid").(string)
	request["cpuCore"] = d.Get("cpu_core").(string)
	request["memory"] = d.Get("memory").(string)
	uuid := d.Id()
	if d.Id() == "" {
		uuid = d.Get("uuid").(string)
	}
	request["uuid"] = uuid
	return request
}

// Action request
// TODO - Documentation
func (vs *instanceUtils) GetInstanceActionRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["action"] = d.Get("action").(string)
	uuid := d.Id()
	if d.Id() == "" {
		uuid = d.Get("uuid").(string)
	}
	request["uuid"] = uuid
	return request
}

// Action request
// TODO - Documentation
func (vs *instanceUtils) GetInstanceIsoActionRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["isoUuid"] = d.Get("iso_uuid").(string)
	request["action"] = d.Get("action").(string)
	uuid := d.Id()
	if d.Id() == "" {
		uuid = d.Get("uuid").(string)
	}
	request["uuid"] = uuid
	return request
}
