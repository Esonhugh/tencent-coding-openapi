package git

type DeleteGitDeployKeyReq struct {
	Action      string `json:"Action"`      // DeleteGitDeployKey
	DeployKeyID int64  `json:"DeployKeyId"` // 1
	DepotID     int64  `json:"DepotId"`     // 1
}

func (req *DeleteGitDeployKeyReq) SetAction() string {
	req.Action = "DeleteGitDeployKey"
	return req.Action
}

type DeleteGitDeployKeyResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 5649c350-f27a-a84a-3cff-f76c36f03280
	} `json:"Response"`
}

// DeleteGitDeployKey 删除部署公钥
func (g *GitClient) DeleteGitDeployKey(req DeleteGitDeployKeyReq) (resp DeleteGitDeployKeyResp, err error) {
	err = g.Call(&req, &resp)
	return
}
