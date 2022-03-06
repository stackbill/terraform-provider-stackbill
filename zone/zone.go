package zone

import "terraform-provider-stackbill/http"

var (
	httpClient  http.HttpClient = http.NewHttpClient()
	zoneApiObj  ZoneApi         = NewZoneApi()
	zoneDataObj ZoneData        = NewZoneData()
)
