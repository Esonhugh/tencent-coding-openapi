package git

type DeleteGitDepotReq struct {
	Action  string `json:"Action"`  // DeleteGitDepot
	DepotID int64  `json:"DepotId"` // 1759
}

func (req *DeleteGitDepotReq) SetAction() string {
	req.Action = "DeleteGitDepot"
	return req.Action
}

type DeleteGitDepotResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 4d7694ed-5818-3ea5-4834-a9d74cb2a594
	} `json:"Response"`
}

// DeleteGitDepot 删除代码仓库
func (g *GitClient) DeleteGitDepot(req DeleteGitDepotReq) (resp DeleteGitDepotResp, err error) {
	err = g.Call(&req, &resp)
	return
}
