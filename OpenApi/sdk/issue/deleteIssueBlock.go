package issue

type DeleteIssueBlockReq struct {
	BlockIssueCode int64  `json:"BlockIssueCode"` // 4
	IssueCode      int64  `json:"IssueCode"`      // 1
	ProjectName    string `json:"ProjectName"`    // project-demo
}

func (req *DeleteIssueBlockReq) SetAction() string {
	return "DeleteIssueBlock"
}

type DeleteIssueBlockResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DeleteIssueBlock 删除前置事项
func (i *IssueClient) DeleteIssueBlock(req DeleteIssueBlockReq) (resp DeleteIssueBlockResp, err error) {
	err = i.Call(&req, &resp)
	return
}
