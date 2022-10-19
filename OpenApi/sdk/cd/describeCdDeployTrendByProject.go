package cd

type DescribeCdDeployTrendByProjectReq struct {
	Action      string `json:"Action"`      // DescribeCdDeployTrendByProject
	EndAt       string `json:"EndAt"`       // 2020-12-03 23:59:59
	ProjectName string `json:"ProjectName"` // test
	StartAt     string `json:"StartAt"`     // 2020-12-01 00:00:00
}

func (req *DescribeCdDeployTrendByProjectReq) SetAction() string {
	req.Action = "DescribeCdDeployTrendByProject"
	return req.Action
}

type DescribeCdDeployTrendByProjectResp struct {
	Response struct {
		Data struct {
			Details []struct {
				Application   string `json:"Application"` // dev
				CdDeployTrend []struct {
					Date         string `json:"Date"`         // 2020-12-01
					SucceedCount int64  `json:"SucceedCount"` // 0
					TotalCount   int64  `json:"TotalCount"`   // 0
				} `json:"CdDeployTrend"`
				SuccessCount int64 `json:"SuccessCount"` // 2
				TotalCount   int64 `json:"TotalCount"`   // 2
			} `json:"Details"`
			EndDate   string `json:"EndDate"`   // 2020-12-03
			StartDate string `json:"StartDate"` // 2020-12-01
			Total     struct {
				CdDeployTrend []struct {
					Date         string `json:"Date"`         // 2020-12-01
					SucceedCount int64  `json:"SucceedCount"` // 0
					TotalCount   int64  `json:"TotalCount"`   // 0
				} `json:"CdDeployTrend"`
				SuccessCount int64 `json:"SuccessCount"` // 3
				TotalCount   int64 `json:"TotalCount"`   // 3
			} `json:"Total"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // b9591c5a-b7d5-8d0a-fa91-5d9b7fc6c88d
	} `json:"Response"`
}

// DescribeCdDeployTrendByProject 根据项目名获取关联应用的发布趋势
func (c *CdClient) DescribeCdDeployTrendByProject(req DescribeCdDeployTrendByProjectReq) (resp DescribeCdDeployTrendByProjectResp, err error) {
	err = c.Call(&req, &resp)
	return
}
