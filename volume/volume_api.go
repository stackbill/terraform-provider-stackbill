package volume

import (
	"errors"
	"terraform-provider-stackbill/api"
)

// New Volume Api
func NewVolumeApi() VolumeApi {
	return &volumeApi{}
}

// Volume api interface
type VolumeApi interface {
	ListVolumes(string, string, interface{}) (string, error)
	VolumeActions(map[string]interface{}, interface{}) (string, error)
	CreateVolume(map[string]interface{}, interface{}) (string, error)
	DeleteVolume(string, interface{}) (string, error)
}

// Volume api object
type volumeApi struct {
}

// List Zones
// TODO - Documentation
func (v *volumeApi) ListVolumes(zoneUuid string, uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetVolumeListApi(zoneUuid)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Actions -  start / Stop / Restart
// TODO - Documentation
func (v *volumeApi) VolumeActions(vr map[string]interface{}, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := ""
	switch vr["action"].(string) {
	case ATTACH:
		endPoint = api.GetVolumeAttachApi(vr["uuid"].(string)) + "&instanceUuid=" + vr["instanceUuid"].(string)
	case DETACH:
		endPoint = api.GetVolumeDetachApi(vr["uuid"].(string))
	default:
		return "", errors.New("Invalid action provided...!")
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}

// List Zones
// TODO - Documentation
func (v *volumeApi) CreateVolume(vc map[string]interface{}, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetVolumeCreateApi()
	response, err := httpClient.PostJson(endPoint, vc)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Delete Volume
// TODO - Documentation
func (v *volumeApi) DeleteVolume(uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetVolumeDeleteApi(uuid)
	response, err := httpClient.Delete(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}
