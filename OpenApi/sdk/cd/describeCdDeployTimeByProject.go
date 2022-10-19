package cd

type DescribeCdDeployTimeByProjectReq struct {
	Action      string `json:"Action"`      // DescribeCdDeployTimeByProject
	EndAt       string `json:"EndAt"`       // 2020-12-31 23:59:59
	ProjectName string `json:"ProjectName"` // test
	StartAt     string `json:"StartAt"`     // 2020-12-01 00:00:00
}

func (req *DescribeCdDeployTimeByProjectReq) SetAction() string {
	req.Action = "DescribeCdDeployTimeByProject"
	return req.Action
}

type DescribeCdDeployTimeByProjectResp struct {
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
		RequestID string `json:"RequestId"` // 88e099d1-487e-cd40-0673-c51e865a9e81
	} `json:"Response"`
}

// DescribeCdDeployTimeByProject 根据项目名获取关联应用的发布时长
func (c *CdClient) DescribeCdDeployTimeByProject(req DescribeCdDeployTimeByProjectReq) (resp DescribeCdDeployTimeByProjectResp, err error) {
	err = c.Call(&req, &resp)
	return
}
