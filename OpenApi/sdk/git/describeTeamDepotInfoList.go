package git

type DescribeTeamDepotInfoListReq struct {
	Action string `json:"Action"` // DescribeTeamDepotInfoList
	TeamID int64  `json:"TeamId"` // 1
	UserID int64  `json:"UserId"` // 1
}

func (req *DescribeTeamDepotInfoListReq) SetAction() string {
	req.Action = "DescribeTeamDepotInfoList"
	return req.Action
}

type DescribeTeamDepotInfoListResp struct {
	Response struct {
		DepotData struct {
			Depots []struct {
				Id        int    `json:"Id"`
				Name      string `json:"Name"`
				HttpsUrl  string `json:"HttpsUrl"`
				ProjectId int    `json:"ProjectId"`
				SshUrl    string `json:"SshUrl"`
				WebUrl    string `json:"WebUrl"`
				VscType   string `json:"VscType"`
			} `json:"Depots"`
		} `json:"DepotData"`
		RequestID string `json:"RequestId"` // ba708820-d6ed-bb70-a628-d25d00a3c1ae
	} `json:"Response"`
}

// DescribeTeamDepotInfoList 获取团队下仓库列表
func (g *GitClient) DescribeTeamDepotInfoList(req DescribeTeamDepotInfoListReq) (resp DescribeTeamDepotInfoListResp, err error) {
	err = g.Call(&req, &resp)
	return
}
