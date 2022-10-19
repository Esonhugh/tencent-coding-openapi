package issue

type CreateIssueBlockReq struct {
	BlockIssueCode int64  `json:"BlockIssueCode"` // 4
	IssueCode      int64  `json:"IssueCode"`      // 1
	ProjectName    string `json:"ProjectName"`    // project-demo
}

func (req *CreateIssueBlockReq) SetAction() string {
	return "CreateIssueBlock"
}

type CreateIssueBlockResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// CreateIssueBlock 添加前置事项
func (i *IssueClient) CreateIssueBlock(req CreateIssueBlockReq) (resp CreateIssueBlockResp, err error) {
	err = i.Call(&req, &resp)
	return
}
