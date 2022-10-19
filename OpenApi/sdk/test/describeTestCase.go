package test

type DescribeTestCaseReq struct {
	Action      string `json:"Action"`             // DescribeTestCase
	CaseID      int64  `json:"CaseId"`             // 1
	ProjectName int64  `json:"ProjectName,string"` // 2
}

func (req *DescribeTestCaseReq) SetAction() string {
	req.Action = "DescribeTestCase"
	return req.Action
}

type DescribeTestCaseResp struct {
	Response struct {
		Data struct {
			Case struct {
				CreatedAt   string `json:"CreatedAt"` // 2020-09-15 16:34:21
				CreatedBy   int64  `json:"CreatedBy"` // 1
				CustomSteps []struct {
					Content  string `json:"Content"`  // expected
					Expected string `json:"Expected"` // steps
				} `json:"CustomSteps"`
				Expected     string `json:"Expected"`     // expected
				ID           int64  `json:"Id"`           // 65
				Preconds     string `json:"Preconds"`     // preconds
				Priority     int64  `json:"Priority"`     // 2
				SectionID    int64  `json:"SectionId"`    // 1
				Sort         int64  `json:"Sort"`         // 15
				Steps        string `json:"Steps"`        // setps
				TemplateType string `json:"TemplateType"` // STEPS
				Title        string `json:"Title"`        // helel
				UpdatedAt    string `json:"UpdatedAt"`    // 2020-09-15 16:34:21
			} `json:"Case"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a0
	} `json:"Response"`
}

// DescribeTestCase 测试用例详情
func (t *TestClient) DescribeTestCase(req DescribeTestCaseReq) (resp DescribeTestCaseResp, err error) {
	err = t.Call(&req, &resp)
	return
}
