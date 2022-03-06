package iso

import "terraform-provider-stackbill/http"

var (
	httpClient http.HttpClient = http.NewHttpClient()
	isoApiObj  IsoApi          = NewIsoApi()
	isoDataObj IsoData         = NewIsoData()
)
