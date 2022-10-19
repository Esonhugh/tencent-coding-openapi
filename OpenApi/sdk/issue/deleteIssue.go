package issue

type DeleteIssueReq struct {
	Action      string `json:"Action"`      // DeleteIssue
	IssueCode   int64  `json:"IssueCode"`   // 3
	ProjectName string `json:"ProjectName"` // demo-project
}

func (req *DeleteIssueReq) SetAction() string {
	req.Action = "DeleteIssue"
	return req.Action
}

type DeleteIssueResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 135f80f0-0577-6421-293e-ec231ff3b337
	} `json:"Response"`
}

// DeleteIssue 删除事项
func (i *IssueClient) DeleteIssue(req DeleteIssueReq) (resp DeleteIssueResp, err error) {
	err = i.Call(&req, &resp)
	return
}
