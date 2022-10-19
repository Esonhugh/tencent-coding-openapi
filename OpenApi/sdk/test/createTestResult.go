package test

type CreateTestResultReq struct {
	Action           string   `json:"Action"`           // CreateTestResult
	CustomStepStatus []string `json:"CustomStepStatus"` // PASSED
	ProjectName      string   `json:"ProjectName"`      // project-demo
	Status           string   `json:"Status"`           // PASSED
	TestID           int64    `json:"TestId"`           // 110
}

func (req *CreateTestResultReq) SetAction() string {
	req.Action = "CreateTestResult"
	return req.Action
}

type CreateTestResultResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a1
	} `json:"Response"`
}

// CreateTestResult 测试任务添加测试结果
func (t *TestClient) CreateTestResult(req CreateTestResultReq) (resp CreateTestResultResp, err error) {
	err = t.Call(&req, &resp)
	return
}
