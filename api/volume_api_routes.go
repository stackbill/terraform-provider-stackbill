package api

var (
	volumeList   = "volume/volumeList"
	attachVolume = "volume/attachVolume"
	detachVolume = "volume/detachVolume"
	createVolume = "volume/createVolume"
	deleteVolume = "volume/deleteVolume"
)

// Volume List api
func GetVolumeListApi(zoneUuid string) string {
	return END_POINT + "/" + volumeList + "?zoneUuid=" + zoneUuid
}

// Volume Attach api
func GetVolumeAttachApi(uuid string) string {
	return END_POINT + "/" + attachVolume + "?uuid=" + uuid
}

// Volume Detach api
func GetVolumeDetachApi(uuid string) string {
	return END_POINT + "/" + detachVolume + "?uuid=" + uuid
}

// Volume Create api
func GetVolumeCreateApi() string {
	return END_POINT + "/" + createVolume
}

// Volume Delete api
func GetVolumeDeleteApi(uuid string) string {
	return END_POINT + "/" + deleteVolume + "/" + uuid
}
