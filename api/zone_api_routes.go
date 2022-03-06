package api

var (
	zoneList = "zone/zonelist"
)

// Zone List api
func GetZoneListApi() string {
	return END_POINT + "/" + zoneList
}
