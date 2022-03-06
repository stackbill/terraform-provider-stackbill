package iso

import (
	"terraform-provider-stackbill/api"
)

// New Iso Api
func NewIsoApi() IsoApi {
	return &isoApi{}
}

// Iso api interface
type IsoApi interface {
	ListIso(string, string, interface{}) (string, error)
}

// Iso api object
type isoApi struct {
}

// List Isos
// TODO - Documentation
func (nt *isoApi) ListIso(zoneUuid string, uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetNewTemplateListApi(zoneUuid)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
