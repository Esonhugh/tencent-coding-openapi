package test

type CreateTestCaseSectionReq struct {
	Action      string `json:"Action"`      // CreateTestCaseSection
	Name        string `json:"Name"`        // 分组名称
	ParentID    int64  `json:"ParentId"`    // 1
	ProjectName string `json:"ProjectName"` // 项目名称
}

func (req *CreateTestCaseSectionReq) SetAction() string {
	req.Action = "CreateTestCaseSection"
	return req.Action
}

type CreateTestCaseSectionResp struct {
	Response struct {
		Data struct {
			Section struct {
				ID       int64 `json:"Id"`          // 4
				Name     int64 `json:"Name,string"` // 2211
				ParentID int64 `json:"ParentId"`    // 0
				Sort     int64 `json:"Sort"`        // 4
			} `json:"Section"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // fe82d911-5911-33f4-26f0-b2287bb8ce4c
	} `json:"Response"`
}

// CreateTestCaseSection 创建测试用例分组
func (t *TestClient) CreateTestCaseSection(req CreateTestCaseSectionReq) (resp CreateTestCaseSectionResp, err error) {
	err = t.Call(&req, &resp)
	return
}
