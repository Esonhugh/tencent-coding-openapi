package issue

type DescribeIssueStatusChangeLogListReq struct {
	Action      string  `json:"Action"`      // DescribeIssueStatusChangeLogList
	IssueCode   []int64 `json:"IssueCode"`   // 380
	ProjectName string  `json:"ProjectName"` // my-demo
}

func (req *DescribeIssueStatusChangeLogListReq) SetAction() string {
	req.Action = "DescribeIssueStatusChangeLogList"
	return req.Action
}

type DescribeIssueStatusChangeLogListResp struct {
	Response struct {
		Logs struct {
			List []struct {
				CreatedAt   int64 `json:"CreatedAt"` // 1.617095165e+12
				IssueCode   int64 `json:"IssueCode"` // 380
				IssueStatus struct {
					CreatedAt   int64  `json:"CreatedAt"`   // 1.551937083e+12
					Description string `json:"Description"` //
					ID          int64  `json:"Id"`          // 244
					Index       int64  `json:"Index"`       // 6
					IsSystem    bool   `json:"IsSystem"`    // true
					Name        string `json:"Name"`        // 未评估
					Type        string `json:"Type"`        // TODO
					UpdatedAt   int64  `json:"UpdatedAt"`   // 1.551937083e+12
				} `json:"IssueStatus"`
				StatusID   int64  `json:"StatusId"`   // 244
				StatusName string `json:"StatusName"` // 未评估
			} `json:"List"`
		} `json:"Logs"`
		RequestID string `json:"RequestId"` // 570322d0-b7bf-28d6-ca17-f8ab645da
	} `json:"Response"`
}

// DescribeIssueStatusChangeLogList 获取事项的状态变更历史
func (i *IssueClient) DescribeIssueStatusChangeLogList(req DescribeIssueStatusChangeLogListReq) (resp DescribeIssueStatusChangeLogListResp, err error) {
	err = i.Call(&req, &resp)
	return
}
