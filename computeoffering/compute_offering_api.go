package computeoffering

import (
	"terraform-provider-stackbill/api"
)

// New Compute Offering Api
func NewComputeOfferingApi() ComputeOfferingApi {
	return &computeOfferingApi{}
}

// Compute Offering Api Interface
type ComputeOfferingApi interface {
	ListComputeOfferings(string, string, interface{}) (string, error)
}

// Compute Offering Api object
type computeOfferingApi struct {
}

/* List Compute offerings
@Param - Zone uuid - required
@Param - Compute offering uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (co *computeOfferingApi) ListComputeOfferings(zoneUuid string, uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetComputeOfferingListApi(zoneUuid)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}
