package git

type ModifyDepotNameReq struct {
	Action    string `json:"Action"`    // ModifyDepotName
	DepotID   int64  `json:"DepotId"`   // 791507
	DepotName string `json:"DepotName"` // demo-update
	UserID    int64  `json:"UserId"`    // 89189
}

func (req *ModifyDepotNameReq) SetAction() string {
	req.Action = "ModifyDepotName"
	return req.Action
}

type ModifyDepotNameResp struct {
	Response struct {
		DepotInfo struct {
			HttpsURL  string `json:"HttpsUrl"`  // https://e.coding.net/codingcorp/demo/demo-update.git
			ID        int64  `json:"Id"`        // 791507
			Name      string `json:"Name"`      // demo-update
			ProjectID int64  `json:"ProjectId"` // 779272
			SshURL    string `json:"SshUrl"`    // git@git.e.coding.net:codingcorp/demo/demo-update.git
			VcsType   string `json:"VcsType"`   // git
			WebURL    string `json:"WebUrl"`    // https://codingcorp.coding.net/p/demo/d/demo-update
		} `json:"DepotInfo"`
		RequestID string `json:"RequestId"` // 212012ad-e391-0f54-4216-5817abda8f3e
	} `json:"Response"`
}

// ModifyDepotName 修改仓库名称
func (g *GitClient) ModifyDepotName(req ModifyDepotNameReq) (resp ModifyDepotNameResp, err error) {
	err = g.Call(&req, &resp)
	return
}
