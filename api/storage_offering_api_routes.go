package api

var (
	storageOfferingList = "storage/storageOfferingList"
)

// Storage Offering List Api
func GetStorageOfferingListApi(zoneUuid string) string {
	return END_POINT + "/" + storageOfferingList + "?zoneUuid=" + zoneUuid
}
