package issue

type DescribeTeamIssueTypeListReq struct {
	Action string `json:"Action"` // DescribeTeamIssueTypeList
}

func (req *DescribeTeamIssueTypeListReq) SetAction() string {
	req.Action = "DescribeTeamIssueTypeList"
	return req.Action
}

type DescribeTeamIssueTypeListResp struct {
	IssueTypes []struct {
		Description string `json:"Description"` //
		ID          int64  `json:"Id"`          // 1
		IsSystem    bool   `json:"IsSystem"`    // false
		IssueType   string `json:"IssueType"`   // REQUIREMENT
		Name        string `json:"Name"`        // 用户故事
	} `json:"IssueTypes"`
}

// DescribeTeamIssueTypeList 查询企业所有事项类型列表
func (i *IssueClient) DescribeTeamIssueTypeList(req DescribeTeamIssueTypeListReq) (resp DescribeTeamIssueTypeListResp, err error) {
	err = i.Call(&req, &resp)
	return
}
