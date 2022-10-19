package iteration

type PlanIterationIssueReq struct {
	Action        string  `json:"Action"`        // PlanIterationIssue
	IssueCode     []int64 `json:"IssueCode"`     // 1
	IterationCode int64   `json:"IterationCode"` // 1
	ProjectName   string  `json:"ProjectName"`   // demo-project
}

func (req *PlanIterationIssueReq) SetAction() string {
	req.Action = "PlanIterationIssue"
	return req.Action
}

type PlanIterationIssueResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// PlanIterationIssue 批量规划迭代
func (i *IterationClient) PlanIterationIssue(req PlanIterationIssueReq) (resp PlanIterationIssueResp, err error) {
	err = i.Call(&req, &resp)
	return
}
