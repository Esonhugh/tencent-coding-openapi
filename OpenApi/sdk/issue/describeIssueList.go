package issue

type DescribeIssueListReq struct {
	Action     string `json:"Action"` // DescribeIssueList
	Conditions []struct {
		Key   string `json:"Key"`   // STATUS_TYPE
		Value string `json:"Value"` // TODO
	} `json:"Conditions"`
	IssueType   string `json:"IssueType"`   // ALL
	Limit       int64  `json:"Limit"`       // 10
	Offset      int64  `json:"Offset"`      // 0
	ProjectName string `json:"ProjectName"` // demo-project
}

func (req *DescribeIssueListReq) SetAction() string {
	req.Action = "DescribeIssueList"
	return req.Action
}

type DescribeIssueListResp struct {
	Response struct {
		IssueList []struct {
			AssigneeID      int64  `json:"AssigneeId"`      // 0
			Code            int64  `json:"Code"`            // 27
			CompletedAt     int64  `json:"CompletedAt"`     // 0
			CreatedAt       int64  `json:"CreatedAt"`       // 1.598843625e+12
			CreatorID       int64  `json:"CreatorId"`       // 1
			Description     string `json:"Description"`     //
			DueDate         int64  `json:"DueDate"`         // 0
			IssueStatusID   int64  `json:"IssueStatusId"`   // 4
			IssueStatusName string `json:"IssueStatusName"` // 未开始
			IssueStatusType string `json:"IssueStatusType"` // TODO
			IssueTypeDetail struct {
				Description string `json:"Description"` //
				ID          int64  `json:"Id"`          // 1
				IsSystem    bool   `json:"IsSystem"`    // false
				IssueType   string `json:"IssueType"`   // REQUIREMENT
				Name        string `json:"Name"`        // 用户故事
			} `json:"IssueTypeDetail"`
			Iteration struct {
				Code   int64  `json:"Code"`   // 1
				Name   string `json:"Name"`   // Q1 Spring1
				Status string `json:"Status"` // PROCESSING
			} `json:"Iteration"`
			IterationID  int64   `json:"IterationId"`       // 1
			Name         string  `json:"Name"`              // this demo issue
			ParentType   string  `json:"ParentType"`        // REQUIREMENT
			Priority     int64   `json:"Priority,string"`   // 3
			StartDate    int64   `json:"StartDate"`         // 0
			StoryPoint   float64 `json:"StoryPoint,string"` // 0.5
			Type         string  `json:"Type"`              // REQUIREMENT
			UpdatedAt    int64   `json:"UpdatedAt"`         // 1.598922564e+12
			WorkingHours int64   `json:"WorkingHours"`      // 0
		} `json:"IssueList"`
		RequestID string `json:"RequestId"` // 35c5b423-d311-d27f-f48c-a2ddf8b1cc59
	} `json:"Response"`
}

// DescribeIssueList 查询事项列表（旧）
func (i *IssueClient) DescribeIssueList(req DescribeIssueListReq) (resp DescribeIssueListResp, err error) {
	err = i.Call(&req, &resp)
	return
}
