package test

type ModifyTestCaseSectionReq struct {
	Action      string `json:"Action"`      // ModifyTestCaseSection
	Name        string `json:"Name"`        // 分组名称
	ParentID    int64  `json:"ParentId"`    // 1
	ProjectName string `json:"ProjectName"` // 项目名称
	SectionID   int64  `json:"SectionId"`   // 1
}

func (req *ModifyTestCaseSectionReq) SetAction() string {
	req.Action = "ModifyTestCaseSection"
	return req.Action
}

type ModifyTestCaseSectionResp struct {
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

// ModifyTestCaseSection 修改测试用例分组
func (t *TestClient) ModifyTestCaseSection(req ModifyTestCaseSectionReq) (resp ModifyTestCaseSectionResp, err error) {
	err = t.Call(&req, &resp)
	return
}
