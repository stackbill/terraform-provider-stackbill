package zone

import (
	"terraform-provider-stackbill/api"
)

// New Zone Api
func NewZoneApi() ZoneApi {
	return &zoneApi{}
}

// Zone api interface
type ZoneApi interface {
	ListZones(string, interface{}) (string, error)
}

// Zone api object
type zoneApi struct {
}

// List Zones
// TODO - Documentation
func (z *zoneApi) ListZones(uuid string, meta interface{}) (string, error) {
	// Meta information
	// m := meta.(*auth.AuthKeys)
	// apiKey := m.ApiKey
	// secretKey := m.SecretKey
	endPoint := api.GetZoneListApi()
	if uuid != "" {
		endPoint += "?uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint)
	if err != nil {
		return "", err
	}
	return response, nil
}
