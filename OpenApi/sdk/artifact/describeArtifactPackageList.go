package artifact

type DescribeArtifactPackageListReq struct {
	Action        string `json:"Action"`        // DescribeArtifactPackageList
	PackagePrefix string `json:"PackagePrefix"` // root/
	PageNumber    int64  `json:"PageNumber"`    // 1
	PageSize      int64  `json:"PageSize"`      // 10
	ProjectID     int64  `json:"ProjectId"`     // 450
	Repository    string `json:"Repository"`    // generic
}

func (req *DescribeArtifactPackageListReq) SetAction() string {
	req.Action = "DescribeArtifactPackageList"
	return req.Action
}

type DescribeArtifactPackageListResp struct {
	Response struct {
		Data struct {
			InstanceSet []struct {
				CreatedAt                  int64   `json:"CreatedAt"`                  // 1.603426362e+09
				Description                string  `json:"Description"`                //
				ID                         int64   `json:"Id"`                         // 354263
				LatestVersionID            int64   `json:"LatestVersionId"`            // 1.589443e+06
				LatestVersionName          float64 `json:"LatestVersionName,string"`   // 1.0
				LatestVersionReleaseStatus int64   `json:"LatestVersionReleaseStatus"` // 1
				Name                       string  `json:"Name"`                       // demo
				ReleaseStrategy            int64   `json:"ReleaseStrategy"`            // 4
				RepoID                     int64   `json:"RepoId"`                     // 200866
				VersionCount               int64   `json:"VersionCount"`               // 1
			} `json:"InstanceSet"`
			PageNumber int64 `json:"PageNumber"` // 1
			PageSize   int64 `json:"PageSize"`   // 10
			TotalCount int64 `json:"TotalCount"` // 1
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeArtifactPackageList 查询制品包（镜像）列表
func (a *ArtifactClient) DescribeArtifactPackageList(req DescribeArtifactPackageListReq) (resp DescribeArtifactPackageListResp, err error) {
	err = a.Call(&req, &resp)
	return
}
