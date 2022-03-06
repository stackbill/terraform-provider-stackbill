package api

var (
	sshKeyList  = "sshkey/sshkeyList"
	resetSshKey = "instance/resetSSHkey"
)

// SshKey List api
func GetSshkeyListApi() string {
	return END_POINT + "/" + sshKeyList
}

// SshKey List api
func GetResetSshkeyApi(uuid string, sshKeyId string) string {
	return END_POINT + "/" + resetSshKey + "?uuid=" + uuid + "&sshUuid=" + sshKeyId
}
