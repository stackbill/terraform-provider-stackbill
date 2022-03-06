package volume

import "terraform-provider-stackbill/http"

var (
	httpClient               http.HttpClient       = http.NewHttpClient()
	volumeApiObj             VolumeApi             = NewVolumeApi()
	volumeDataObj            VolumeData            = NewVolumeData()
	volumeUtilObj            VolumeUtils           = NewVolumeUtils()
	volumeActionsResourceObj VolumeActionsResource = NewVolumeActionsResource()
	volumeResourceObj        VolumeResource        = NewVolumeResource()
)

// Volume Actions
const (
	ATTACH = "Attach"
	DETACH = "Detach"
)
