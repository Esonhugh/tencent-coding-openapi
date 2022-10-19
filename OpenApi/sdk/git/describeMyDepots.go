package git

type DescribeMyDepotsReq struct {
	Action string `json:"Action"` // DescribeMyDepots
}

func (req *DescribeMyDepotsReq) SetAction() string {
	req.Action = "DescribeMyDepots"
	return req.Action
}

type DescribeMyDepotsResp struct {
	Response struct {
		Payload struct {
			Depots []struct {
				HttpsURL  string `json:"HttpsUrl"`  // xx
				ID        int64  `json:"Id"`        // 1
				Name      string `json:"Name"`      // xx
				ProjectID int64  `json:"ProjectId"` // 1
				SshURL    string `json:"SshUrl"`    // xx
				VcsType   string `json:"VcsType"`   // git
				WebURL    string `json:"WebUrl"`    // xx
			} `json:"Depots"`
		} `json:"Payload"`
		RequestID string `json:"RequestId"` // 4bc86b5d-03ea-44a6-ae7b-b1239e92eea4
	} `json:"Response"`
}

// DescribeMyDepots 查询当前用户的仓库列表
func (g *GitClient) DescribeMyDepots(req DescribeMyDepotsReq) (resp DescribeMyDepotsResp, err error) {
	err = g.Call(&req, &resp)
	return
}
