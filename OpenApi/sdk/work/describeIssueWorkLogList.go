package work

type DescribeIssueWorkLogListReq struct {
	IssueCode   int64  `json:"IssueCode"`   // 1
	ProjectName string `json:"ProjectName"` // project-demo
}

func (req *DescribeIssueWorkLogListReq) SetAction() string {
	return "DescribeIssueWorkLogList"
}

type DescribeIssueWorkLogListResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
		WorkLogs  []struct {
			CreatedAt      int64  `json:"CreatedAt"`      // 1.608689693852e+12
			ID             int64  `json:"Id"`             // 1
			IssueID        int64  `json:"IssueId"`        // 1
			RecordHours    int64  `json:"RecordHours"`    // 1
			RemainingHours int64  `json:"RemainingHours"` // 2
			UserID         int64  `json:"UserId"`         // 1
			WorkingDesc    string `json:"WorkingDesc"`    // xx
		} `json:"WorkLogs"`
	} `json:"Response"`
}

// DescribeIssueWorkLogList 查询工时日志列表
func (w *WorkClient) DescribeIssueWorkLogList(req DescribeIssueWorkLogListReq) (resp DescribeIssueWorkLogListResp, err error) {
	err = w.Call(&req, &resp)
	return
}
