package git

type DescribeProjectDepotInfoListReq struct {
	Action    string `json:"Action"`    // DescribeProjectDepotInfoList
	ProjectID int64  `json:"ProjectId"` // 234
}

func (req *DescribeProjectDepotInfoListReq) SetAction() string {
	req.Action = "DescribeProjectDepotInfoList"
	return req.Action
}

type DescribeProjectDepotInfoListResp struct {
	Response struct {
		DepotData struct {
			Depots []struct {
				Description string `json:"Description"` // foo
				HttpsURL    string `json:"HttpsUrl"`    // https://e.coding.net/demo/demo-project/depot-1.git
				ID          int64  `json:"Id"`          // 1
				Name        string `json:"Name"`        // depot-1
				SshURL      string `json:"SshUrl"`      // git@e.coding.net:demo/demo-project/depot-1.git
				VcsType     string `json:"VcsType"`     // git
			} `json:"Depots"`
		} `json:"DepotData"`
		RequestID string `json:"RequestId"` // e1d658b4-7e89-febe-2072-e68a9d9852e3
	} `json:"Response"`
}

// DescribeProjectDepotInfoList 查询项目下仓库信息列表
func (g *GitClient) DescribeProjectDepotInfoList(req DescribeProjectDepotInfoListReq) (resp DescribeProjectDepotInfoListResp, err error) {
	err = g.Call(&req, &resp)
	return
}
