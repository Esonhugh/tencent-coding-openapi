package cd

type DescribeCdDeployTrendByApplicationsReq struct {
	Action      string   `json:"Action"`      // DescribeCdDeployTrendByApplications
	Application []string `json:"Application"` // dev
	EndAt       string   `json:"EndAt"`       // 2020-12-03 23:59:59
	StartAt     string   `json:"StartAt"`     // 2020-12-01 00:00:00
}

func (req *DescribeCdDeployTrendByApplicationsReq) SetAction() string {
	req.Action = "DescribeCdDeployTrendByApplications"
	return req.Action
}

type DescribeCdDeployTrendByApplicationsResp struct {
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
		RequestID string `json:"RequestId"` // 2d9211a8-07d7-927c-3e80-7d78cf793ffe
	} `json:"Response"`
}

// DescribeCdDeployTrendByApplications 根据应用名列表获取发布趋势
func (c *CdClient) DescribeCdDeployTrendByApplications(req DescribeCdDeployTrendByApplicationsReq) (resp DescribeCdDeployTrendByApplicationsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
