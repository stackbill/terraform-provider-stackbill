package snapshot

import "terraform-provider-stackbill/http"

var (
	httpClient            http.HttpClient    = http.NewHttpClient()
	vmSnapshotApiObj      VmSnapshotApi      = NewVmSnapshotApi()
	vmSnapshotUtilsObj    VmSnapshotUtils    = NewVmSnapshotUtils()
	vmSnapshotResourceObj VmSnapshotResource = NewVmSnapshotResource()
	vmSnapshotDataObj     VmSnapshotData     = NewVmSnapshotData()
)
