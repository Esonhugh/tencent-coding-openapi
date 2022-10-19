package test

type DeleteTestCaseSectionReq struct {
	Action      string `json:"Action"`             // DeleteTestCaseSection
	ProjectName int64  `json:"ProjectName,string"` // 2
	SectionID   int64  `json:"SectionId"`          // 1
}

func (req *DeleteTestCaseSectionReq) SetAction() string {
	req.Action = "DeleteTestCaseSection"
	return req.Action
}

type DeleteTestCaseSectionResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a1
	} `json:"Response"`
}

// DeleteTestCaseSection 删除测试用例分组
func (t *TestClient) DeleteTestCaseSection(req DeleteTestCaseSectionReq) (resp DeleteTestCaseSectionResp, err error) {
	err = t.Call(&req, &resp)
	return
}
