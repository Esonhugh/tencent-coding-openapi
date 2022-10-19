package test

type DescribeTestReq struct {
	Action      string `json:"Action"`      // DescribeTest
	ProjectName string `json:"ProjectName"` // xx
	TestID      int64  `json:"TestId"`      // 1
}

func (req *DescribeTestReq) SetAction() string {
	req.Action = "DescribeTest"
	return req.Action
}

type DescribeTestResp struct {
	Response struct {
		Data struct {
			Test struct {
				AssignedToID int64 `json:"AssignedToId"` // 0
				Case         struct {
					Attachments []struct {
						CreatedAt string `json:"CreatedAt"` // 2021-07-09 11:04:19
						ID        int64  `json:"Id"`        // 14
						Name      string `json:"Name"`      // filename
						URL       string `json:"Url"`       // https://coding.com
					} `json:"Attachments"`
					CreatedAt    string        `json:"CreatedAt"`    // 2021-05-24 15:08:58
					CreatedBy    int64         `json:"CreatedBy"`    // 1
					CustomSteps  []interface{} `json:"CustomSteps"`  // <nil>
					Expected     string        `json:"Expected"`     // 提示密码不能为空
					ID           int64         `json:"Id"`           // 1
					Preconds     string        `json:"Preconds"`     // 用户手机号，验证码和密码
					Priority     int64         `json:"Priority"`     // 2
					SectionID    int64         `json:"SectionId"`    // 2
					Sort         int64         `json:"Sort"`         // 1
					Steps        string        `json:"Steps"`        // 输入手机号和验证码
					TemplateType string        `json:"TemplateType"` // TEXT
					Title        string        `json:"Title"`        // 手机注册
					UpdatedAt    string        `json:"UpdatedAt"`    // 2021-07-09 11:04:21
				} `json:"Case"`
				CaseID      int64    `json:"CaseId"`      // 1
				ID          int64    `json:"Id"`          // 1533
				IsCompleted bool     `json:"IsCompleted"` // false
				Priority    int64    `json:"Priority"`    // 2
				SectionID   int64    `json:"SectionId"`   // 2
				SectionPath []string `json:"SectionPath"` // 分组一级
				Sort        int64    `json:"Sort"`        // 1
				Status      string   `json:"Status"`      // PASSED
				TestedAt    string   `json:"TestedAt"`    // 2021-07-09 11:39:24
				TestedBy    int64    `json:"TestedBy"`    // 1
				Title       string   `json:"Title"`       // 手机注册
			} `json:"Test"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 2efcd378-f6cf-83c0-b705-86896740b45e
	} `json:"Response"`
}

// DescribeTest 测试任务详情
func (t *TestClient) DescribeTest(req DescribeTestReq) (resp DescribeTestResp, err error) {
	err = t.Call(&req, &resp)
	return
}
