package issue

type DescribeSubIssueListReq struct {
	Action      string `json:"Action"`      // DescribeSubIssueList
	IssueCode   int64  `json:"IssueCode"`   // 1
	ProjectName string `json:"ProjectName"` // project-demo
}

func (req *DescribeSubIssueListReq) SetAction() string {
	req.Action = "DescribeSubIssueList"
	return req.Action
}

type DescribeSubIssueListResp struct {
	Response struct {
		RequestID    int64 `json:"RequestId,string"` // 1
		SubIssueList []struct {
			AssigneeID      int64   `json:"AssigneeId"`        // 0
			Code            int64   `json:"Code"`              // 27
			CompletedAt     int64   `json:"CompletedAt"`       // 0
			CreatedAt       int64   `json:"CreatedAt"`         // 1.598843625e+12
			CreatorID       int64   `json:"CreatorId"`         // 1
			Description     string  `json:"Description"`       //
			DueDate         int64   `json:"DueDate"`           // 0
			IssueStatusID   int64   `json:"IssueStatusId"`     // 4
			IssueStatusName string  `json:"IssueStatusName"`   // 未开始
			IssueStatusType string  `json:"IssueStatusType"`   // TODO
			IterationID     int64   `json:"IterationId"`       // 0
			Name            string  `json:"Name"`              // xxx
			ParentType      string  `json:"ParentType"`        // REQUIREMENT
			Priority        int64   `json:"Priority,string"`   // 3
			StartDate       int64   `json:"StartDate"`         // 0
			StoryPoint      float64 `json:"StoryPoint,string"` // 0.5
			Type            string  `json:"Type"`              // REQUIREMENT
			UpdatedAt       int64   `json:"UpdatedAt"`         // 1.598922564e+12
			WorkingHours    int64   `json:"WorkingHours"`      // 0
		} `json:"SubIssueList"`
	} `json:"Response"`
}

// DescribeSubIssueList 查询子事项列表
func (i *IssueClient) DescribeSubIssueList(req DescribeSubIssueListReq) (resp DescribeSubIssueListResp, err error) {
	err = i.Call(&req, &resp)
	return
}
