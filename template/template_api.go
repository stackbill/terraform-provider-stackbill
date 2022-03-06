package template

import (
	"terraform-provider-stackbill/api"
)

// New Template Api
func NewTemplateApi() TemplateApi {
	return &templateApi{}
}

// Template api interface
type TemplateApi interface {
	ListTemplates(string, string, interface{}) (string, error)
}

// Template api object
type templateApi struct {
}

// List Templates
// TODO - Documentation
func (nt *templateApi) ListTemplates(zoneUuid string, uuid string, meta interface{}) (string, error) {
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
	return response, nil
}
