package work

type DeleteIssueWorkHoursReq struct {
	IssueCode              int64  `json:"IssueCode"`              // 1
	ProjectName            string `json:"ProjectName"`            // project-demo
	RollbackRemainingHours bool   `json:"RollbackRemainingHours"` // true
	WorkHourLogID          int64  `json:"WorkHourLogId"`          // 2
}

func (req *DeleteIssueWorkHoursReq) SetAction() string {
	return "DeleteIssueWorkHours"
}

type DeleteIssueWorkHoursResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // xx
	} `json:"Response"`
}

// DeleteIssueWorkHours 删除工时日志
func (w *WorkClient) DeleteIssueWorkHours(req DeleteIssueWorkHoursReq) (resp DeleteIssueWorkHoursResp, err error) {
	err = w.Call(&req, &resp)
	return
}
