package test

type CreateTestDefectReq struct {
	Action      string `json:"Action"`      // CreateTestDefect
	DefectID    int64  `json:"DefectId"`    // 1
	ProjectName string `json:"ProjectName"` // xx
	TestID      int64  `json:"TestId"`      // 1
}

func (req *CreateTestDefectReq) SetAction() string {
	req.Action = "CreateTestDefect"
	return req.Action
}

type CreateTestDefectResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a1
	} `json:"Response"`
}

// CreateTestDefect 测试任务关联缺陷
func (t *TestClient) CreateTestDefect(req CreateTestDefectReq) (resp CreateTestDefectResp, err error) {
	err = t.Call(&req, &resp)
	return
}
