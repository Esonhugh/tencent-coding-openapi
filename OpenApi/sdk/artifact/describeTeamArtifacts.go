package artifact

type DescribeTeamArtifactsReq struct {
	PageNumber int64 `json:"PageNumber"` // 1
	PageSize   int64 `json:"PageSize"`   // 10
	Rule       struct {
		ArtifactType []int64 `json:"ArtifactType"` // 1
		Package      []struct {
			Algorithm string `json:"Algorithm"`    // EQUAL
			Value     int64  `json:"Value,string"` // 1
		} `json:"Package"`
		PackageVersion []struct {
			Algorithm string `json:"Algorithm"` // EQUAL
			Value     string `json:"Value"`     // latest
		} `json:"PackageVersion"`
		ProjectName []string `json:"ProjectName"` // test
		Repository  []string `json:"Repository"`  // g1
	} `json:"Rule"`
}

func (req *DescribeTeamArtifactsReq) SetAction() string {
	return "DescribeTeamArtifacts"
}

type DescribeTeamArtifactsResp struct {
	Response struct {
		Data struct {
			InstanceSet []struct {
				ArtifactType   int64   `json:"ArtifactType"`          // 1
				CreatedAt      int64   `json:"CreatedAt,string"`      // 1629873218
				Description    string  `json:"Description"`           //
				DownloadCount  int64   `json:"DownloadCount"`         // 0
				Hash           string  `json:"Hash"`                  // 12ce92814df58fdc16863d65a8dee59c0a7ccf10
				Package        int64   `json:"Package,string"`        // 1
				PackageVersion int64   `json:"PackageVersion,string"` // 1
				PkgID          int64   `json:"PkgId"`                 // 22
				ProjectID      int64   `json:"ProjectId"`             // 453
				ProjectName    string  `json:"ProjectName"`           // test
				ReleaseStatus  int64   `json:"ReleaseStatus"`         // 2
				RepoID         int64   `json:"RepoId"`                // 3
				Repository     string  `json:"Repository"`            // g1
				Size           float64 `json:"Size"`                  // 1.430511474609375e-05
				VersionID      int64   `json:"VersionId"`             // 29
			} `json:"InstanceSet"`
			PageNumber int64 `json:"PageNumber"` // 1
			PageSize   int64 `json:"PageSize"`   // 10
			TotalCount int64 `json:"TotalCount"` // 1
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeTeamArtifacts 查询团队下制品版本列表
func (a *ArtifactClient) DescribeTeamArtifacts(req DescribeTeamArtifactsReq) (resp DescribeTeamArtifactsResp, err error) {
	err = a.Call(&req, &resp)
	return
}
