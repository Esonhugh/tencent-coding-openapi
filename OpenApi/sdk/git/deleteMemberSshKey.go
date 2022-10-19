package git

type DeleteMemberSshKeyReq struct {
	Action       string `json:"Action"`       // DeleteMemberSshKey
	MemberUserID int64  `json:"MemberUserId"` // 2
	SshKeyID     int64  `json:"SshKeyId"`     // 2
}

func (req *DeleteMemberSshKeyReq) SetAction() string {
	req.Action = "DeleteMemberSshKey"
	return req.Action
}

type DeleteMemberSshKeyResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 13c2cc53-88fe-ad64-6253-865d7ccb69ed
	} `json:"Response"`
}

// DeleteMemberSshKey 删除团队成员的 SSH 公钥
func (g *GitClient) DeleteMemberSshKey(req DeleteMemberSshKeyReq) (resp DeleteMemberSshKeyResp, err error) {
	err = g.Call(&req, &resp)
	return
}
