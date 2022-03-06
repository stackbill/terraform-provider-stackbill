package instance

import "terraform-provider-stackbill/http"

var (
	httpClient                     http.HttpClient             = http.NewHttpClient()
	instanceApiObj                 InstanceApi                 = NewInstanceApi()
	instanceUtilsObj               InstanceUtils               = NewInstanceUtils()
	instanceActionsResourceObj     InstanceActionsResource     = NewInstanceActionsResource()
	instanceIsoActionsResourceObj  InstanceIsoActionsResource  = NewInstanceIsoActionsResource()
	instanceDataObj                InstanceData                = NewInstanceData()
	instanceObj                    InstanceResource            = NewInstance()
	instanceUpdateNameResourceObj  InstanceUpdateNameResource  = NewInstanceUpdaeNameResource()
	instanceResizeResourceObj      InstanceResizeResource      = NewInstanceResizeResource()
	instanceResetSshkeyResourceObj InstanceResetSshkeyResource = NewInstanceResetSshKeyResource()
)

const (
	START = "Start"
	STOP  = "Stop"
)

const (
	ATTACH = "Attach"
	DETACH = "Detach"
)
