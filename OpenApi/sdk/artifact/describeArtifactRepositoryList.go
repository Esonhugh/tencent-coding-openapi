package artifact

type DescribeArtifactRepositoryListReq struct {
	Action     string `json:"Action"`     // DescribeArtifactRepositoryList
	PageNumber int64  `json:"PageNumber"` // 1
	PageSize   int64  `json:"PageSize"`   // 1
	ProjectID  int64  `json:"ProjectId"`  // 450
	Type       int64  `json:"Type"`       // 1
}

func (req *DescribeArtifactRepositoryListReq) SetAction() string {
	req.Action = "DescribeArtifactRepositoryList"
	return req.Action
}

type DescribeArtifactRepositoryListResp struct {
	Response struct {
		Data struct {
			InstanceSet []struct {
				AccessLevel     int64  `json:"AccessLevel"`     // 1
				CreatedAt       int64  `json:"CreatedAt"`       // 1.605600358e+09
				Description     string `json:"Description"`     //
				ID              int64  `json:"Id"`              // 246823
				Name            int64  `json:"Name,string"`     // 1232
				ProjectID       int64  `json:"ProjectId"`       // 7.936e+06
				ReleaseStrategy int64  `json:"ReleaseStrategy"` // 1
				TeamID          int64  `json:"TeamId"`          // 64
				Type            int64  `json:"Type"`            // 1
			} `json:"InstanceSet"`
			PageNumber int64 `json:"PageNumber"` // 1
			PageSize   int64 `json:"PageSize"`   // 10
			TotalCount int64 `json:"TotalCount"` // 3
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 3c140219-cfe9-470e-b241-907877d6fb03
	} `json:"Response"`
}

// DescribeArtifactRepositoryList 查询制品仓库列表
func (a *ArtifactClient) DescribeArtifactRepositoryList(req DescribeArtifactRepositoryListReq) (resp DescribeArtifactRepositoryListResp, err error) {
	err = a.Call(&req, &resp)
	return
}
