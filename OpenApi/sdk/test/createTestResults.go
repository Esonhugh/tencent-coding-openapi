package test

type CreateTestResultsReq struct {
	Action      string  `json:"Action"`      // CreateTestResults
	CaseIDs     []int64 `json:"CaseIds"`     // 1
	ProjectName string  `json:"ProjectName"` // xx
	RunID       int64   `json:"RunId"`       // 1
	Status      string  `json:"Status"`      // xx
}

func (req *CreateTestResultsReq) SetAction() string {
	req.Action = "CreateTestResults"
	return req.Action
}

type CreateTestResultsResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 980cbd88-caac-626a-8c8d-8019acf41bb4
	} `json:"Response"`
}

// CreateTestResults 测试任务状态批量更新
func (t *TestClient) CreateTestResults(req CreateTestResultsReq) (resp CreateTestResultsResp, err error) {
	err = t.Call(&req, &resp)
	return
}
