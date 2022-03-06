package network

import "terraform-provider-stackbill/http"

var (
	httpClient                http.HttpClient        = http.NewHttpClient()
	networkApiObj             NetworkApi             = NewNetworkApi()
	networkDataObj            NetworkData            = NewNetworkData()
	networkOfferingDataObj    NetworkOfferingData    = NewNetworkOfferingData()
	vpcNetworkOfferingDataObj VpcNetworkOfferingData = NewVpcNetworkOfferingData()
	networkUtilObj            NetworkUtils           = NewNetworkUtils()
	networkResourceObj        NetworkResource        = NewNetworkResource()
	networkActionResourceObj  NetworkActionResource  = NewNetworkActionResource()
)

const (
	ADD    = "Add"
	DELETE = "Delete"
)
