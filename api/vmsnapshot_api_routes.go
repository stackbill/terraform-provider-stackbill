package api

var (
	vmSnapShotList   = "vmsnapshot/vmsnapshotList"
	vmSnapshot       = "vmsnapshot/createVmSnapshot"
	vmSnapshotDelete = "vmsnapshot/deleteVmSnapshot"
)

// Instance Snapshot api
func GetVmSnapshotCreateApi() string {
	return END_POINT + "/" + vmSnapshot
}

// Snapshot List api
func GetVmSnapshotListApi(zoneUuid string) string {
	return END_POINT + "/" + vmSnapShotList + "?zoneUuid=" + zoneUuid
}

// Zone List api
func GetVmSnapshotDeleteApi(uuid string) string {
	return END_POINT + "/" + vmSnapshotDelete + "/" + uuid
}
