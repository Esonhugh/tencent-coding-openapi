package git

type DescribeSshKeyReq struct {
	Action   string `json:"Action"`   // DescribeSshKey
	SshKeyID int64  `json:"SshKeyId"` // 2
}

func (req *DescribeSshKeyReq) SetAction() string {
	req.Action = "DescribeSshKey"
	return req.Action
}

type DescribeSshKeyResp struct {
	Response struct {
		RequestID   string `json:"RequestId"` // 27440043-d529-7e27-3209-ad69157d486e
		SshKeyInfos []struct {
			CreatedAt      int64  `json:"CreatedAt"`      // 1.63480837e+12
			ExpirationDate string `json:"ExpirationDate"` // 2022-09-09
			FingerPrint    string `json:"FingerPrint"`    // 0c:36:f5:60:fd:da:3b:be:22:49:38:96:64:50:49:9f
			HasExpired     bool   `json:"HasExpired"`     // false
			ID             int64  `json:"Id"`             // 2
			OwnerID        int64  `json:"OwnerId"`        // 44
			Title          string `json:"Title"`          // Hello
		} `json:"SshKeyInfos"`
	} `json:"Response"`
}

// DescribeSshKey 获取当前用户 SSH 公钥
func (g *GitClient) DescribeSshKey(req DescribeSshKeyReq) (resp DescribeSshKeyResp, err error) {
	err = g.Call(&req, &resp)
	return
}
