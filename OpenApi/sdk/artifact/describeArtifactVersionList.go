package artifact

type DescribeArtifactVersionListReq struct {
	Action     string `json:"Action"`     // DescribeArtifactVersionList
	Package    string `json:"Package"`    // Pro
	PageNumber int64  `json:"PageNumber"` // 1
	PageSize   int64  `json:"PageSize"`   // 10
	ProjectID  int64  `json:"ProjectId"`  // 450
	Repository string `json:"Repository"` // generic
}

func (req *DescribeArtifactVersionListReq) SetAction() string {
	req.Action = "DescribeArtifactVersionList"
	return req.Action
}

type DescribeArtifactVersionListResp struct {
	Response struct {
		Data struct {
			InstanceSet []struct {
				CreatedAt     int64   `json:"CreatedAt"`      // 1.603426362e+09
				Description   string  `json:"Description"`    //
				DownloadCount int64   `json:"DownloadCount"`  // 0
				Hash          string  `json:"Hash"`           // sha1 514854de2a43d015ed21042ac953058241cbfeb7
				ID            int64   `json:"Id"`             // 1.589443e+06
				PkgID         int64   `json:"PkgId"`          // 354263
				ReleaseStatus int64   `json:"ReleaseStatus"`  // 0
				Version       float64 `json:"Version,string"` // 1.0
			} `json:"InstanceSet"`
			PageNumber int64 `json:"PageNumber"` // 1
			PageSize   int64 `json:"PageSize"`   // 10
			TotalCount int64 `json:"TotalCount"` // 1
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeArtifactVersionList 查询制品版本列表
func (a *ArtifactClient) DescribeArtifactVersionList(req DescribeArtifactVersionListReq) (resp DescribeArtifactVersionListResp, err error) {
	err = a.Call(&req, &resp)
	return
}
