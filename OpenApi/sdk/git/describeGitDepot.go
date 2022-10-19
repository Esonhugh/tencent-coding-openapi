package git

type DescribeGitDepotReq struct {
	Action  string `json:"Action"`  // DescribeGitDepot
	DepotID int64  `json:"DepotId"` // 8.521598e+06
}

func (req *DescribeGitDepotReq) SetAction() string {
	req.Action = "DescribeGitDepot"
	return req.Action
}

type DescribeGitDepotResp struct {
	Response struct {
		Depot struct {
			HttpsURL  string `json:"HttpsUrl"`  // https://e.coding.net/codingcorp/pual_test/demo.git
			ID        int64  `json:"Id"`        // 8.521598e+06
			Name      string `json:"Name"`      // demo
			ProjectID int64  `json:"ProjectId"` // 8.104562e+06
			SshURL    string `json:"SshUrl"`    // git@e.coding.net:codingcorp/pual_test/demo.git
			VcsType   string `json:"VcsType"`   // git
			WebURL    string `json:"WebUrl"`    // https://codingcorp.coding.net/p/pual_test/d/demo
		} `json:"Depot"`
		RequestID string `json:"RequestId"` // 92f7ae17-cb38-d8c5-5144-0a76b6a7ee7b
	} `json:"Response"`
}

// DescribeGitDepot 获取代码仓库详情
func (g *GitClient) DescribeGitDepot(req DescribeGitDepotReq) (resp DescribeGitDepotResp, err error) {
	err = g.Call(&req, &resp)
	return
}
