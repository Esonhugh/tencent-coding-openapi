package test

type DescribeRequirementTestCaseListReq struct {
	Action      string `json:"Action"`      // DescribeRequirementTestCaseList
	IssueID     int64  `json:"IssueId"`     // 1
	ProjectName string `json:"ProjectName"` // xx
}

func (req *DescribeRequirementTestCaseListReq) SetAction() string {
	req.Action = "DescribeRequirementTestCaseList"
	return req.Action
}

type DescribeRequirementTestCaseListResp struct {
	Response struct {
		Data struct {
			Cases []struct {
				CreatedAt   string `json:"CreatedAt"` // 2020-09-15 15:32:59
				CreatedBy   int64  `json:"CreatedBy"` // 1
				CustomSteps []struct {
					Content  string `json:"Content"`  // expected
					Expected string `json:"Expected"` // steps
				} `json:"CustomSteps"`
				Expected     interface{} `json:"Expected"`     // <nil>
				ID           int64       `json:"Id"`           // 53
				Preconds     string      `json:"Preconds"`     // preconds
				Priority     int64       `json:"Priority"`     // 2
				SectionID    int64       `json:"SectionId"`    // 1
				Sort         int64       `json:"Sort"`         // 3
				Steps        interface{} `json:"Steps"`        // <nil>
				TemplateType string      `json:"TemplateType"` // STEPS
				Title        string      `json:"Title"`        // helel22
				UpdatedAt    string      `json:"UpdatedAt"`    // 2020-09-15 15:35:28
			} `json:"Cases"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // b18a050b-7530-a56b-cb37-98831ef3cf6c
	} `json:"Response"`
}

// DescribeRequirementTestCaseList 需求关联的测试用例列表
func (t *TestClient) DescribeRequirementTestCaseList(req DescribeRequirementTestCaseListReq) (resp DescribeRequirementTestCaseListResp, err error) {
	err = t.Call(&req, &resp)
	return
}
