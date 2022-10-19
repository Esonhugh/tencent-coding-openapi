package git

type DeleteTeamLevelDepotSpecReq struct {
	Action        string `json:"Action"`        // DeleteTeamLevelDepotSpec
	DepotSpecName string `json:"DepotSpecName"` // depot-spec-name
}

func (req *DeleteTeamLevelDepotSpecReq) SetAction() string {
	req.Action = "DeleteTeamLevelDepotSpec"
	return req.Action
}

type DeleteTeamLevelDepotSpecResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 00b6cddb-bd7d-964e-6885-45bbc387bd64
	} `json:"Response"`
}

// DeleteTeamLevelDepotSpec 删除团队级别的分支规范
func (g *GitClient) DeleteTeamLevelDepotSpec(req DeleteTeamLevelDepotSpecReq) (resp DeleteTeamLevelDepotSpecResp, err error) {
	err = g.Call(&req, &resp)
	return
}
