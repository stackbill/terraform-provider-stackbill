package template

import "terraform-provider-stackbill/http"

var (
	httpClient      http.HttpClient = http.NewHttpClient()
	templateApiObj  TemplateApi     = NewTemplateApi()
	templateDataObj TemplateData    = NewTemplateData()
)
