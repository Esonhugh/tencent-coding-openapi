package git

type DescribeMemberSshKeyReq struct {
	Action       string `json:"Action"`       // DescribeMemberSshKey
	MemberUserID int64  `json:"MemberUserId"` // 2
	SshKeyID     int64  `json:"SshKeyId"`     // 2
}

func (req *DescribeMemberSshKeyReq) SetAction() string {
	req.Action = "DescribeMemberSshKey"
	return req.Action
}

type DescribeMemberSshKeyResp struct {
	Response struct {
		RequestID   string `json:"RequestId"` // d16367d7-6377-5361-e16b-c8efe0b463df
		SshKeyInfos []struct {
			CreatedAt      int64  `json:"CreatedAt"`      // 1.642398559986e+12
			ExpirationDate string `json:"ExpirationDate"` // 2025-12-31
			FingerPrint    string `json:"FingerPrint"`    // 32:a9:dd:de:4c:3f:df:6f:0a:43:23:48:23:06:d2:31
			HasExpired     bool   `json:"HasExpired"`     // false
			ID             int64  `json:"Id"`             // 2
			OwnerID        int64  `json:"OwnerId"`        // 2
			Title          string `json:"Title"`          // create a ssh key for member
		} `json:"SshKeyInfos"`
	} `json:"Response"`
}

// DescribeMemberSshKey 查询团队成员的 SSH 公钥列表
func (g *GitClient) DescribeMemberSshKey(req DescribeMemberSshKeyReq) (resp DescribeMemberSshKeyResp, err error) {
	err = g.Call(&req, &resp)
	return
}
