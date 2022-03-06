package network

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Network util Object
func NewNetworkUtils() NetworkUtils {
	return &networkUtils{}
}

//	Network Interface
type NetworkUtils interface {
	GetCreateNetworkRequest(d *schema.ResourceData) map[string]interface{}
	GetCreateNetworkOptionalRequest(d *schema.ResourceData) map[string]interface{}
	GetNetworkActionRequest(d *schema.ResourceData) map[string]interface{}
}

// Network utils Object
type networkUtils struct {
}

// Snapshot Request
// TODO - Documentation
func (nt *networkUtils) GetCreateNetworkRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["name"] = d.Get("name").(string)
	request["networkOfferingUuid"] = d.Get("network_offering_uuid").(string)
	if d.Get("virtual_machine_uuid").(string) != "" {
		request["virtualmachineUuid"] = d.Get("virtual_machine_uuid").(string)
	}
	request["zoneUuid"] = d.Get("zone_uuid").(string)
	return request
}

// Snapshot Request
// TODO - Documentation
func (nt *networkUtils) GetCreateNetworkOptionalRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["isPublic"] = d.Get("is_public").(bool)
	request["securityGroupId"] = d.Get("security_group_id").(string)
	return request
}

// Action Request
// TODO - Documentation
func (nt *networkUtils) GetNetworkActionRequest(d *schema.ResourceData) map[string]interface{} {
	request := make(map[string]interface{})
	request["networkUuid"] = d.Get("network_uuid").(string)
	request["uuid"] = d.Get("virutal_machine_uuid").(string)
	return request
}
