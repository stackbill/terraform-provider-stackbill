package storageoffering

import (
	"terraform-provider-stackbill/api"
)

// New Storage Offering Api
func NewStorageOfferingApi() StorageOfferingApi {
	return &storageOfferingApi{}
}

// Storage Offering Api Interfae
type StorageOfferingApi interface {
	ListStorageOfferings(string, string, interface{}) (string, error)
}

// Storage Offering Api object
type storageOfferingApi struct {
}

/* List Storage offerings
@Param - Zone uuid - required
@Param - Storage offering uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (co *storageOfferingApi) ListStorageOfferings(zoneId string, uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetStorageOfferingListApi(zoneId)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}
