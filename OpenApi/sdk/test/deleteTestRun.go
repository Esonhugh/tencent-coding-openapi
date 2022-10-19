package test

type DeleteTestRunReq struct {
	Action      string `json:"Action"`      // DeleteTestRun
	ProjectName string `json:"ProjectName"` // project name
	RunID       int64  `json:"RunId"`       // 1
}

func (req *DeleteTestRunReq) SetAction() string {
	req.Action = "DeleteTestRun"
	return req.Action
}

type DeleteTestRunResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // daba6d8a-0a62-e6fd-f502-a4594888561a
	} `json:"Response"`
}

// DeleteTestRun 删除测试计划
func (t *TestClient) DeleteTestRun(req DeleteTestRunReq) (resp DeleteTestRunResp, err error) {
	err = t.Call(&req, &resp)
	return
}
