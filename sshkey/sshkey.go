package sshkey

import "terraform-provider-stackbill/http"

var (
	httpClient    http.HttpClient = http.NewHttpClient()
	sshkeyApiObj  SshkeyApi       = NewSshkeyApi()
	sshkeyDataObj SshkeyData      = NewSshkeyData()
)
