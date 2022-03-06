package instance

import (
	"encoding/json"
	"errors"
	"terraform-provider-stackbill/api"
)

// Instance util Object
func NewInstanceApi() InstanceApi {
	return &instanceApi{}
}

//Instance Interfae
type InstanceApi interface {
	CreateInstance(map[string]interface{}, interface{}) (string, error)
	UpdateInstanceName(map[string]interface{}, interface{}) (string, error)
	ResetSshKey(map[string]interface{}, interface{}) (string, error)
	ResizeInstance(map[string]interface{}, interface{}) (string, error)
	InstanceActions(map[string]interface{}, interface{}) (string, error)
	InstanceIsoActions(map[string]interface{}, interface{}) (string, error)
	ListInstances(string, string, interface{}) (string, error)
	DeleteInstance(string, interface{}) (string, error)
	GetInstanceStatus(string, interface{}) (string, error)
}

//Instance utils Object
type instanceApi struct {
}

// Update Instance Name
// TODO - Documentation
func (ia *instanceApi) CreateInstance(cr map[string]interface{}, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetInstanceCreateApi()
	response, err := httpClient.PostJson(endPoint, cr)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Update Instance Name
// TODO - Documentation
func (ia *instanceApi) ResetSshKey(rsr map[string]interface{}, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetResetSshkeyApi(rsr["uuid"].(string), rsr["sshKeyId"].(string))
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Reset ssh key
// TODO - Documentation
func (ia *instanceApi) UpdateInstanceName(ur map[string]interface{}, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetInstanceNameUpdateApi()
	response, err := httpClient.PutJson(endPoint, ur)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Resize VM
// TODO - Documentation
func (vs *instanceApi) ResizeInstance(r map[string]interface{}, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	qryString := "?uuid=" + r["uuid"].(string) + "&offeringUuid=" + r["offeringUuid"].(string)
	qryString += "&memory=" + r["memory"].(string) + "&cpuCore=" + r["cpuCore"].(string)
	endPoint := api.GetInstanceResizeApi() + qryString
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Actions -  start / Stop / Restart
// TODO - Documentation
func (vs *instanceApi) InstanceActions(ar map[string]interface{}, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := ""
	switch ar["action"].(string) {
	case START:
		endPoint = api.GetInstanceStartApi() + "?uuid=" + ar["uuid"].(string)
	case STOP:
		endPoint = api.GetInstanceStopApi() + "?uuid=" + ar["uuid"].(string) + "&forceStop=true"
	default:
		return "", errors.New("Invalid action provided...!")
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Actions -  Attach / Detach
// TODO - Documentation
func (vs *instanceApi) InstanceIsoActions(ar map[string]interface{}, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := ""
	switch ar["action"].(string) {
	case ATTACH:
		endPoint = api.GetInstanceAttachIsoApi(ar["uuid"].(string), ar["isoUuid"].(string))
	case DETACH:
		endPoint = api.GetInstanceDetachIsoApi(ar["uuid"].(string), ar["isoUuid"].(string))
	default:
		return "", errors.New("Invalid action provided...!")
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Delete Instance
// TODO - Documentation
func (vs *instanceApi) DeleteInstance(uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	qryString := "?uuid=" + uuid + "&expunge=true"
	endPoint := api.GetInstanceDeleteApi() + qryString
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Delete Instance
// TODO - Documentation
func (vs *instanceApi) ListInstances(zoneUuid string, uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetInstanceListApi(zoneUuid)
	if uuid != "" {
		endPoint += "&vmUuid=" + uuid
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Resize VM
// TODO - Documentation
func (vs *instanceApi) GetInstanceStatus(uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetInstanceStatusApi(uuid)
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	var resJson map[string]interface{}
	if err := json.Unmarshal([]byte(response), &resJson); err != nil {
		return "", err
	}
	status := resJson["status"].(string)
	return status, nil
}
