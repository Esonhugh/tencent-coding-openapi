package test

type DeleteTestCaseReq struct {
	Action      string `json:"Action"`             // DeleteTestCase
	CaseID      int64  `json:"CaseId"`             // 1
	ProjectName int64  `json:"ProjectName,string"` // 2
}

func (req *DeleteTestCaseReq) SetAction() string {
	req.Action = "DeleteTestCase"
	return req.Action
}

type DeleteTestCaseResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a0
	} `json:"Response"`
}

// DeleteTestCase 删除测试用例
func (t *TestClient) DeleteTestCase(req DeleteTestCaseReq) (resp DeleteTestCaseResp, err error) {
	err = t.Call(&req, &resp)
	return
}
