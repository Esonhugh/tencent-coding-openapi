package git

type ModifyGitDepotUnarchiveReq struct {
	Action  string `json:"Action"`  // ModifyGitDepotUnarchive
	DepotID int64  `json:"DepotId"` // 2
}

func (req *ModifyGitDepotUnarchiveReq) SetAction() string {
	req.Action = "ModifyGitDepotUnarchive"
	return req.Action
}

type ModifyGitDepotUnarchiveResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 87b9311d-23fc-9f49-e16d-fa890f613ebb
	} `json:"Response"`
}

// ModifyGitDepotUnarchive 仓库解除归档
func (g *GitClient) ModifyGitDepotUnarchive(req ModifyGitDepotUnarchiveReq) (resp ModifyGitDepotUnarchiveResp, err error) {
	err = g.Call(&req, &resp)
	return
}
