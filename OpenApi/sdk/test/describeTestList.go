package test

type DescribeTestListReq struct {
	Action      string `json:"Action"`      // DescribeTestList
	Priority    int64  `json:"Priority"`    // 1
	ProjectName string `json:"ProjectName"` // xx
	RunID       int64  `json:"RunId"`       // 1
	Status      string `json:"Status"`      // xx
}

func (req *DescribeTestListReq) SetAction() string {
	req.Action = "DescribeTestList"
	return req.Action
}

type DescribeTestListResp struct {
	Response struct {
		Data struct {
			Tests []struct {
				AssignedToID int64  `json:"AssignedToId"` // 0
				CaseID       int64  `json:"CaseId"`       // 52
				ID           int64  `json:"Id"`           // 71
				IsCompleted  bool   `json:"IsCompleted"`  // true
				Priority     int64  `json:"Priority"`     // 2
				SectionID    int64  `json:"SectionId"`    // 5
				Sort         int64  `json:"Sort"`         // 2
				Status       string `json:"Status"`       // UNTESTED
				TestedAt     string `json:"TestedAt"`     // xx
				Title        string `json:"Title"`        // helel
			} `json:"Tests"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 2efcd378-f6cf-83c0-b705-86896740b45e
	} `json:"Response"`
}

// DescribeTestList 测试任务列表
func (t *TestClient) DescribeTestList(req DescribeTestListReq) (resp DescribeTestListResp, err error) {
	err = t.Call(&req, &resp)
	return
}
