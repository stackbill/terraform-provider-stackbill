package api

import "strconv"

var (
	networkList            = "network/networkList"
	networkOfferingList    = "networkoffering/networkOfferingList"
	vpcNetworkOfferingList = "networkoffering/vpcNetworkOfferingList"
	networkCreate          = "network/createNetwork"
	networkDelete          = "network/deleteNetwork"
	vmNetworkAdd           = "instance/attachNetwork"
	vmNetworkDelete        = "instance/detachNetwork"
)

// Network List api
func GetNewNetworkListApi(zoneUuid string) string {
	return END_POINT + "/" + networkList + "?zoneUuid=" + zoneUuid
}

// Network Offering List api
func GetNewNetworkOfferingListApi(zoneUuid string) string {
	return END_POINT + "/" + networkOfferingList + "?zoneUuid=" + zoneUuid
}

// Network Offering List api
func GetNewVpcNetworkOfferingListApi() string {
	return END_POINT + "/" + vpcNetworkOfferingList
}

// Network Offering List api
func GetNetworkCreateApi(isPublic bool, securityGroupId string) string {
	secrityGroup := ""
	if securityGroupId != "" {
		secrityGroup = "&securityGroupId=" + securityGroupId
	}
	return END_POINT + "/" + networkCreate + "?isPublic=" + strconv.FormatBool(isPublic) + secrityGroup
}

// Network Offering List api
func GetNetworkDeleteApi(uuid string) string {
	return END_POINT + "/" + networkDelete + "/" + uuid
}

// Network Attach api
func GetVmNetworkAddApi() string {
	return END_POINT + "/" + vmNetworkAdd
}

// Network Attach api
func GetVmNetworkDeleteApi() string {
	return END_POINT + "/" + vmNetworkDelete
}
