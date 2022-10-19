package git

type DeleteSshKeyReq struct {
	Action   string `json:"Action"`   // DeleteSshKey
	SshKeyID int64  `json:"SshKeyId"` // 15
}

func (req *DeleteSshKeyReq) SetAction() string {
	req.Action = "DeleteSshKey"
	return req.Action
}

type DeleteSshKeyResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ef532935-ca07-87c8-3ebd-9d89b7aae330
	} `json:"Response"`
}

// DeleteSshKey 删除当前用户的 SSH 公钥
func (g *GitClient) DeleteSshKey(req DeleteSshKeyReq) (resp DeleteSshKeyResp, err error) {
	err = g.Call(&req, &resp)
	return
}
