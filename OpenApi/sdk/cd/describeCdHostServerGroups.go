package cd

type DescribeCdHostServerGroupsReq struct {
	Action     string `json:"Action"`     // DescribeCdHostServerGroups
	Keyword    string `json:"Keyword"`    //
	PageNumber int64  `json:"PageNumber"` // 1
	PageSize   int64  `json:"PageSize"`   // 10
}

func (req *DescribeCdHostServerGroupsReq) SetAction() string {
	req.Action = "DescribeCdHostServerGroups"
	return req.Action
}

type DescribeCdHostServerGroupsResp struct {
	Response struct {
		Data struct {
			HostServerGroups []struct {
				Account      string `json:"Account"` // cred-16-18
				AgentMachine struct {
					ID     int64  `json:"Id"`     // 16
					Name   string `json:"Name"`   // master
					Status string `json:"Status"` // SUCCESS
				} `json:"AgentMachine"`
				DisplayName string `json:"DisplayName"` // local-server-group
				ID          int64  `json:"Id"`          // 18
			} `json:"HostServerGroups"`
			PageNumber int64 `json:"PageNumber"` // 1
			PageSize   int64 `json:"PageSize"`   // 10
			TotalPage  int64 `json:"TotalPage"`  // 1
			TotalRow   int64 `json:"TotalRow"`   // 1
		} `json:"Data"`
		RequestID string `json:"RequestId"` // fe2689f4-72e0-02cf-8f98-e807c393d206
	} `json:"Response"`
}

// DescribeCdHostServerGroups 获取 CD 主机组列表
func (c *CdClient) DescribeCdHostServerGroups(req DescribeCdHostServerGroupsReq) (resp DescribeCdHostServerGroupsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
