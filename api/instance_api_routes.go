package api

var (
	instancePostApi       = "instance/createInstance"
	instanceNameUpdateApi = "instance/updateInstanceName"
	instanceStart         = "instance/startInstance"
	instanceStop          = "instance/stopInstance"
	instanceDelete        = "instance/destroyInstance"
	instanceResize        = "instance/resizeVm"
	instanceList          = "instance/instanceList"
	instanceAttachIso     = "instance/attachIso"
	instanceDetachIso     = "instance/detachIso"
	instanceStatus        = "instance/vmStatus"
)

// Instance create api
func GetInstanceCreateApi() string {
	return END_POINT + "/" + instancePostApi
}

// Instance name update api
func GetInstanceNameUpdateApi() string {
	return END_POINT + "/" + instanceNameUpdateApi
}

// Start instance api
func GetInstanceStartApi() string {
	return END_POINT + "/" + instanceStart
}

// Stop instance api
func GetInstanceStopApi() string {
	return END_POINT + "/" + instanceStop
}

// Delete instance api
func GetInstanceDeleteApi() string {
	return END_POINT + "/" + instanceDelete
}

// Instance Resize api
func GetInstanceResizeApi() string {
	return END_POINT + "/" + instanceResize
}

// Instance List api
func GetInstanceListApi(zoneId string) string {
	return END_POINT + "/" + instanceList + "?zoneUuid=" + zoneId
}

// Instace Attach Iso
func GetInstanceAttachIsoApi(uuid string, isoUUid string) string {
	return END_POINT + "/" + instanceAttachIso + "?uuid=" + uuid + "&isoUuid=" + isoUUid
}

// Instace Detach Iso
func GetInstanceDetachIsoApi(uuid string, isoUUid string) string {
	return END_POINT + "/" + instanceDetachIso + "?uuid=" + uuid + "&isoUuid=" + isoUUid
}

// Instance status api
func GetInstanceStatusApi(uuid string) string {
	return END_POINT + "/" + instanceStatus + "?uuid=" + uuid
}
