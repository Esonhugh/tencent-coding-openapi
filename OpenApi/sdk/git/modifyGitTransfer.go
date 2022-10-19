package git

type ModifyGitTransferReq struct {
	Action          string `json:"Action"`          // ModifyGitTransfer
	DepotID         int64  `json:"DepotId"`         // 1
	TargetProjectID int64  `json:"TargetProjectId"` // 2
}

func (req *ModifyGitTransferReq) SetAction() string {
	req.Action = "ModifyGitTransfer"
	return req.Action
}

type ModifyGitTransferResp struct {
	Response struct {
		NewDepotPath string `json:"NewDepotPath"` // /p/project-e/d/ttt/git
		RequestID    string `json:"RequestId"`    // ad93218b-244b-6055-e073-17df43663f9e
	} `json:"Response"`
}

// ModifyGitTransfer 仓库转移至同团队下的其他项目中
func (g *GitClient) ModifyGitTransfer(req ModifyGitTransferReq) (resp ModifyGitTransferResp, err error) {
	err = g.Call(&req, &resp)
	return
}
