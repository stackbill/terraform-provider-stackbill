package securitygroup

import "terraform-provider-stackbill/http"

var (
	httpClient           http.HttpClient   = http.NewHttpClient()
	securityGroupApiObj  SecurityGroupApi  = NewSecurityGroupApi()
	securityGroupDataObj SecuritygroupData = NewSecuritygroupData()
)
