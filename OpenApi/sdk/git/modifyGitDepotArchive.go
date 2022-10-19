package git

type ModifyGitDepotArchiveReq struct {
	Action  string `json:"Action"`  // ModifyGitDepotArchive
	DepotID int64  `json:"DepotId"` // 2
}

func (req *ModifyGitDepotArchiveReq) SetAction() string {
	req.Action = "ModifyGitDepotArchive"
	return req.Action
}

type ModifyGitDepotArchiveResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 0f569eca-7575-e6fe-d02a-8d67d8c8b6be
	} `json:"Response"`
}

// ModifyGitDepotArchive 仓库归档
func (g *GitClient) ModifyGitDepotArchive(req ModifyGitDepotArchiveReq) (resp ModifyGitDepotArchiveResp, err error) {
	err = g.Call(&req, &resp)
	return
}
