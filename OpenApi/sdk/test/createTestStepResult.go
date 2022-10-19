package test

type CreateTestStepResultReq struct {
	Action      string `json:"Action"`      // CreateTestStepResult
	Actual      string `json:"Actual"`      // Actual
	ProjectName string `json:"ProjectName"` // demo-project
	Status      string `json:"Status"`      // PASSED
	StepIndex   int64  `json:"StepIndex"`   // 2
	StepStatus  string `json:"StepStatus"`  // FAILED
	TestID      int64  `json:"TestId"`      // 2
}

func (req *CreateTestStepResultReq) SetAction() string {
	req.Action = "CreateTestStepResult"
	return req.Action
}

type CreateTestStepResultResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a1
	} `json:"Response"`
}

// CreateTestStepResult 测试任务单独添加某步骤的测试结果
func (t *TestClient) CreateTestStepResult(req CreateTestStepResultReq) (resp CreateTestStepResultResp, err error) {
	err = t.Call(&req, &resp)
	return
}
