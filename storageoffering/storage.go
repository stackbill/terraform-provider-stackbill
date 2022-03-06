package storageoffering

import "terraform-provider-stackbill/http"

var (
	httpClient             http.HttpClient     = http.NewHttpClient()
	storageOfferingApiObj  StorageOfferingApi  = NewStorageOfferingApi()
	storageOfferingDataObj StorageOfferingData = NewStorageOfferingData()
)
