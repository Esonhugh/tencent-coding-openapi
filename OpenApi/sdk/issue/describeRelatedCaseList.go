package issue

type DescribeRelatedCaseListReq struct {
	Action      string `json:"Action"`      // DescribeRelatedCaseList
	IssueCode   int64  `json:"IssueCode"`   // 1
	ProjectName string `json:"ProjectName"` // my-demo
}

func (req *DescribeRelatedCaseListReq) SetAction() string {
	req.Action = "DescribeRelatedCaseList"
	return req.Action
}

type DescribeRelatedCaseListResp struct {
	Response struct {
		Cases []struct {
			ID       int64  `json:"Id"`       // 1
			Name     string `json:"Name"`     // 注册商城验证码验证通过
			Priority int64  `json:"Priority"` // 1
		} `json:"Cases"`
		RequestID string `json:"RequestId"` // b1410375-b707-6415-85cc-f0c5ec78a5ba
	} `json:"Response"`
}

// DescribeRelatedCaseList 事项关联的测试用例
func (i *IssueClient) DescribeRelatedCaseList(req DescribeRelatedCaseListReq) (resp DescribeRelatedCaseListResp, err error) {
	err = i.Call(&req, &resp)
	return
}
