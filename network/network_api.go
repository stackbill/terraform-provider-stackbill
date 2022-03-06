package network

import (
	"errors"
	"terraform-provider-stackbill/api"
)

// New Network Api
func NewNetworkApi() NetworkApi {
	return &networkApi{}
}

// Network api interface
type NetworkApi interface {
	ListNetworks(string, string, interface{}) (string, error)
	ListNetworkOfferings(string, string, interface{}) (string, error)
	ListVpcNetworkOfferings(string, interface{}) (string, error)
	CreateNetwork(map[string]interface{}, map[string]interface{}, interface{}) (string, error)
	DeleteNetwork(string, interface{}) (string, error)
	NetworkActions(string, map[string]interface{}, interface{}) (string, error)
}

// Network api object
type networkApi struct {
}

/* List Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) ListNetworks(zoneUuid string, uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetNewNetworkListApi(zoneUuid)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

/* List Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) ListNetworkOfferings(zoneUuid string, uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetNewNetworkOfferingListApi(zoneUuid)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

/* List Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) ListVpcNetworkOfferings(uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetNewVpcNetworkOfferingListApi()
	if uuid != "" {
		endPoint += "?uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

/* Create Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) CreateNetwork(cn map[string]interface{}, co map[string]interface{}, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetNetworkCreateApi(co["isPublic"].(bool), co["securityGroupId"].(string))
	response, err := httpClient.PostJson(endPoint, cn)
	if err != nil {
		return "", err
	}
	return response, nil
}

/* Delete Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) DeleteNetwork(uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetNetworkDeleteApi(uuid)
	response, err := httpClient.Delete(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Actions -  start / Stop / Restart
// TODO - Documentation
func (nt *networkApi) NetworkActions(action string, nr map[string]interface{}, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := ""
	switch action {
	case ADD:
		endPoint = api.GetVmNetworkAddApi()
	case DELETE:
		endPoint = api.GetVmNetworkDeleteApi()
	default:
		return "", errors.New("Invalid action provided...!")
	}
	response, err := httpClient.PostJson(endPoint, nr)
	if err != nil {
		return "", err
	}
	return response, nil
}
