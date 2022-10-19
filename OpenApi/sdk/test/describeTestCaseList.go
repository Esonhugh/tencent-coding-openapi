package test

type DescribeTestCaseListReq struct {
	Action       string `json:"Action"`       // DescribeTestCaseList
	Keyword      string `json:"Keyword"`      // xx
	Priority     int64  `json:"Priority"`     // 1
	ProjectName  string `json:"ProjectName"`  // xx
	SectionID    int64  `json:"SectionId"`    // 1
	TemplateType string `json:"TemplateType"` // xx
}

func (req *DescribeTestCaseListReq) SetAction() string {
	req.Action = "DescribeTestCaseList"
	return req.Action
}

type DescribeTestCaseListResp struct {
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

// DescribeTestCaseList 测试用例列表
func (t *TestClient) DescribeTestCaseList(req DescribeTestCaseListReq) (resp DescribeTestCaseListResp, err error) {
	err = t.Call(&req, &resp)
	return
}
