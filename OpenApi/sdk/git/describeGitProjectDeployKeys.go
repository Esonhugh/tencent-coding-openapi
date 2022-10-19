package git

type DescribeGitProjectDeployKeysReq struct {
	Action    string `json:"Action"`    // DescribeGitProjectDeployKeys
	ProjectID int64  `json:"ProjectId"` // 2
}

func (req *DescribeGitProjectDeployKeysReq) SetAction() string {
	req.Action = "DescribeGitProjectDeployKeys"
	return req.Action
}

type DescribeGitProjectDeployKeysResp struct {
	Response struct {
		Keys []struct {
			AllowWrite     bool   `json:"AllowWrite"`     // true
			CreatedAt      int64  `json:"CreatedAt"`      // 1.642143691e+12
			DepotID        int64  `json:"DepotId"`        // 2
			ExpirationDate string `json:"ExpirationDate"` // 2022-12-12
			FingerPrint    string `json:"FingerPrint"`    // a2:db:f8:eb:56:90:b0:b6:d9:96:0d:99:0c:ba:1a:c5
			HasExpired     bool   `json:"HasExpired"`     // false
			KeyID          int64  `json:"KeyId"`          // 1
			OwnerName      string `json:"OwnerName"`      // coding
			Title          string `json:"Title"`          // deploy key 1
		} `json:"Keys"`
		RequestID string `json:"RequestId"` // 2b2cfdd3-e569-bf07-4281-33c37c428270
	} `json:"Response"`
}

// DescribeGitProjectDeployKeys 查询某项目下的部署公钥列表
func (g *GitClient) DescribeGitProjectDeployKeys(req DescribeGitProjectDeployKeysReq) (resp DescribeGitProjectDeployKeysResp, err error) {
	err = g.Call(&req, &resp)
	return
}
