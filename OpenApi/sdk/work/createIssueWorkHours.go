package work

type CreateIssueWorkHoursReq struct {
	IssueCode     int64  `json:"IssueCode"`      // 1
	ProjectName   string `json:"ProjectName"`    // project-demo
	RemainingHour int64  `json:"RemainingHour"`  // 1
	SpendHour     int64  `json:"SpendHour"`      // 2
	StartAt       int64  `json:"StartAt,string"` // 1608630452740
	WorkingDesc   string `json:"WorkingDesc"`    // xxx
}

func (req *CreateIssueWorkHoursReq) SetAction() string {
	return "CreateIssueWorkHours"
}

type CreateIssueWorkHoursResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// CreateIssueWorkHours 登记工时
func (w *WorkClient) CreateIssueWorkHours(req CreateIssueWorkHoursReq) (resp CreateIssueWorkHoursResp, err error) {
	err = w.Call(&req, &resp)
	return
}
