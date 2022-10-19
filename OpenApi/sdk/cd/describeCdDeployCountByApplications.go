package cd

type DescribeCdDeployCountByApplicationsReq struct {
	Action      string   `json:"Action"`      // DescribeCdDeployCountByApplications
	Application []string `json:"Application"` // dev
	EndAt       string   `json:"EndAt"`       // 2020-12-03 23:59:59
	StartAt     string   `json:"StartAt"`     // 2020-12-01 00:00:00
}

func (req *DescribeCdDeployCountByApplicationsReq) SetAction() string {
	req.Action = "DescribeCdDeployCountByApplications"
	return req.Action
}

type DescribeCdDeployCountByApplicationsResp struct {
	Response struct {
		Data struct {
			Details []struct {
				Application   string `json:"Application"` // dev
				CdDeployCount struct {
					SucceedCount int64 `json:"SucceedCount"` // 2
					TotalCount   int64 `json:"TotalCount"`   // 2
				} `json:"CdDeployCount"`
			} `json:"Details"`
			EndDate   string `json:"EndDate"`   // 2020-12-03
			StartDate string `json:"StartDate"` // 2020-12-01
			Total     struct {
				SucceedCount int64 `json:"SucceedCount"` // 3
				TotalCount   int64 `json:"TotalCount"`   // 3
			} `json:"Total"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 0fb1159d-e76b-4baa-4bd5-2649325d524f
	} `json:"Response"`
}

// DescribeCdDeployCountByApplications 根据应用名列表获取发布次数
func (c *CdClient) DescribeCdDeployCountByApplications(req DescribeCdDeployCountByApplicationsReq) (resp DescribeCdDeployCountByApplicationsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
