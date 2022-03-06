package computeoffering

import "terraform-provider-stackbill/http"

var (
	httpClient             http.HttpClient     = http.NewHttpClient()
	computeOfferingApiObj  ComputeOfferingApi  = NewComputeOfferingApi()
	computeOfferingDataObj ComputeOfferingData = NewComputeOfferingData()
)
