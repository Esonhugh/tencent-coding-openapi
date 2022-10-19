package git

type DescribeDepotByNameInfoReq struct {
	Action      string `json:"Action"`      // DescribeDepotByNameInfo
	DepotName   string `json:"DepotName"`   // depot-2
	ProjectName string `json:"ProjectName"` // project-d
	TeamGk      string `json:"TeamGk"`      // codingcorp
}

func (req *DescribeDepotByNameInfoReq) SetAction() string {
	req.Action = "DescribeDepotByNameInfo"
	return req.Action
}

type DescribeDepotByNameInfoResp struct {
	Response struct {
		Depot struct {
			Description string `json:"Description"` //
			HttpsURL    string `json:"HttpsUrl"`    // http://git.nh108omhq.dev.coding.io/codingcorp/project-d/depot-2.git
			ID          int64  `json:"Id"`          // 2
			Name        string `json:"Name"`        // depot-2
			ProjectID   int64  `json:"ProjectId"`   // 1
			ProjectName string `json:"ProjectName"` // project-d
			SshURL      string `json:"SshUrl"`      // git@git.nh108omhq.dev.coding.io:codingcorp/project-d/depot-2.git
			VcsType     string `json:"VcsType"`     // git
			WebURL      string `json:"WebUrl"`      // http://codingcorp.nh108omhq.dev.coding.io/p/project-d/d/depot-2
		} `json:"Depot"`
		RequestID string `json:"RequestId"` // 5a656ed9-f9e7-4841-9494-7cdaf8e56644
	} `json:"Response"`
}

// DescribeDepotByNameInfo 通过仓库名查询仓库信息
func (g *GitClient) DescribeDepotByNameInfo(req DescribeDepotByNameInfoReq) (resp DescribeDepotByNameInfoResp, err error) {
	err = g.Call(&req, &resp)
	return
}
