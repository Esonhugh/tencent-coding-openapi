package test

type DescribeReportListReq struct {
	Action      string `json:"Action"`      // DescribeReportList
	ProjectName string `json:"ProjectName"` // project name
}

func (req *DescribeReportListReq) SetAction() string {
	req.Action = "DescribeReportList"
	return req.Action
}

type DescribeReportListResp struct {
	Response struct {
		Data struct {
			Reports []struct {
				CreatedAt           string      `json:"CreatedAt"`           // 2021-06-28 15:33:42
				CreatedBy           int64       `json:"CreatedBy"`           // 1
				ID                  int64       `json:"Id"`                  // 1750
				Name                string      `json:"Name"`                // 报告2
				StatisticsEndTime   string      `json:"StatisticsEndTime"`   // 2021-06-28 15:33:42
				StatisticsStartTime string      `json:"StatisticsStartTime"` // 2021-06-16 10:44:26
				Status              string      `json:"Status"`              // UNAVAILABLE
				Summary             interface{} `json:"Summary"`             // <nil>
			} `json:"Reports"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 980cbd88-caac-626a-8c8d-8019acf41bb4
	} `json:"Response"`
}

// DescribeReportList 测试报告列表
func (t *TestClient) DescribeReportList(req DescribeReportListReq) (resp DescribeReportListResp, err error) {
	err = t.Call(&req, &resp)
	return
}
