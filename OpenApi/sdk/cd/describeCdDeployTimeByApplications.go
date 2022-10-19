package cd

type DescribeCdDeployTimeByApplicationsReq struct {
	Action      string   `json:"Action"`      // DescribeCdDeployTimeByApplications
	Application []string `json:"Application"` // dev
	EndAt       string   `json:"EndAt"`       // 2020-12-31 23:59:59
	StartAt     string   `json:"StartAt"`     // 2020-12-01 00:00:00
}

func (req *DescribeCdDeployTimeByApplicationsReq) SetAction() string {
	req.Action = "DescribeCdDeployTimeByApplications"
	return req.Action
}

type DescribeCdDeployTimeByApplicationsResp struct {
	Response struct {
		Data struct {
			Details []struct {
				Application  string `json:"Application"` // dev
				CdDeployTime struct {
					DeployTime int64 `json:"DeployTime"` // 6
					TotalCount int64 `json:"TotalCount"` // 2
				} `json:"CdDeployTime"`
			} `json:"Details"`
			EndDate   string `json:"EndDate"`   // 2020-12-31
			StartDate string `json:"StartDate"` // 2020-12-01
			Total     struct {
				DeployTime int64 `json:"DeployTime"` // 36
				TotalCount int64 `json:"TotalCount"` // 3
			} `json:"Total"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // fc75f5e4-ef08-4469-bb75-9591b85a298d
	} `json:"Response"`
}

// DescribeCdDeployTimeByApplications 根据应用名列表获取发布时长
func (c *CdClient) DescribeCdDeployTimeByApplications(req DescribeCdDeployTimeByApplicationsReq) (resp DescribeCdDeployTimeByApplicationsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
