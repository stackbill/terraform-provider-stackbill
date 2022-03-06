package sshkey

import (
	"terraform-provider-stackbill/api"
)

// New Sshkey Api
func NewSshkeyApi() SshkeyApi {
	return &networkApi{}
}

// Sshkey api interface
type SshkeyApi interface {
	ListSshkeys(interface{}) (string, error)
}

// Sshkey api object
type networkApi struct {
}

/* List Sshkeys
@Param - Zone uuid - required
@Param - Sshkey uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) ListSshkeys(meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetSshkeyListApi()
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}
