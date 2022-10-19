package issue

type DescribeProjectIssueStatusListReq struct {
	Action      string `json:"Action"`      // DescribeProjectIssueStatusList
	IssueType   string `json:"IssueType"`   // REQUIREMENT
	ProjectName string `json:"ProjectName"` // demo-project
}

func (req *DescribeProjectIssueStatusListReq) SetAction() string {
	req.Action = "DescribeProjectIssueStatusList"
	return req.Action
}

type DescribeProjectIssueStatusListResp struct {
	Response struct {
		ProjectIssueStatusList []struct {
			CreatedAt   int64 `json:"CreatedAt"` // 1.5972834e+12
			IsDefault   bool  `json:"IsDefault"` // true
			IssueStatus struct {
				CreatedAt   int64  `json:"CreatedAt"`   // 1.597283396e+12
				Description string `json:"Description"` //
				ID          int64  `json:"Id"`          // 4
				Index       int64  `json:"Index"`       // 3
				IsSystem    bool   `json:"IsSystem"`    // true
				Name        string `json:"Name"`        // 未开始
				Type        string `json:"Type"`        // TODO
				UpdatedAt   int64  `json:"UpdatedAt"`   // 1.597283396e+12
			} `json:"IssueStatus"`
			IssueStatusID int64  `json:"IssueStatusId"` // 4
			IssueType     string `json:"IssueType"`     // REQUIREMENT
			Sort          int64  `json:"Sort"`          // 0
			UpdatedAt     int64  `json:"UpdatedAt"`     // 1.5972834e+12
		} `json:"ProjectIssueStatusList"`
		RequestID string `json:"RequestId"` // dc827006-32db-a74f-eeae-13bec31c8b92
	} `json:"Response"`
}

// DescribeProjectIssueStatusList 查询状态设置
func (i *IssueClient) DescribeProjectIssueStatusList(req DescribeProjectIssueStatusListReq) (resp DescribeProjectIssueStatusListResp, err error) {
	err = i.Call(&req, &resp)
	return
}
