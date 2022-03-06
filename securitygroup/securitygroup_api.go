package securitygroup

import (
	"terraform-provider-stackbill/api"
)

// New Securitygroup Api
func NewSecurityGroupApi() SecurityGroupApi {
	return &securityGroupApi{}
}

// Securitygroup api interface
type SecurityGroupApi interface {
	ListSecurityGroups(string, interface{}) (string, error)
}

// SecurityGroupApi object
type securityGroupApi struct {
}

/* List Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (sg *securityGroupApi) ListSecurityGroups(uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetSecurityGroupListApi()
	if uuid != "" {
		endPoint += "?uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}
