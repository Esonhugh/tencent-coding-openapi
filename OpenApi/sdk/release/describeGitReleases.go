package release

type DescribeGitReleasesReq struct {
	Action     string `json:"Action"`     // DescribeGitReleases
	DepotID    int64  `json:"DepotId"`    // 809883
	FromDate   string `json:"FromDate"`   // 2021-08-31
	PageNumber int64  `json:"PageNumber"` // 1
	PageSize   int64  `json:"PageSize"`   // 10
	Status     int64  `json:"Status"`     // 0
	TagName    string `json:"TagName"`    // TagName11
	ToDate     string `json:"ToDate"`     // 2021-09-02
}

func (req *DescribeGitReleasesReq) SetAction() string {
	req.Action = "DescribeGitReleases"
	return req.Action
}

type DescribeGitReleasesResp struct {
	Response struct {
		ReleasePageList struct {
			Releases []struct {
				Body            int64  `json:"Body,string"`     // 4
				CommitSha       string `json:"CommitSha"`       // 6efc794f0d694fed6cd033ddc984026887af258a
				CreatedAt       int64  `json:"CreatedAt"`       // 1.630488978e+12
				CreatorID       int64  `json:"CreatorId"`       // 1
				DepotID         int64  `json:"DepotId"`         // 809883
				Html            string `json:"Html"`            // <p>4</p>
				ID              int64  `json:"Id"`              // 31586
				Iid             int64  `json:"Iid"`             // 18
				Pre             bool   `json:"Pre"`             // true
				ProjectID       int64  `json:"ProjectId"`       // 797199
				TagName         string `json:"TagName"`         // TagName11
				TargetCommitish string `json:"TargetCommitish"` // 6efc794f0d694fed6cd033ddc984026887af258a
				Title           string `json:"Title"`           // sss
				UpdatedAt       int64  `json:"UpdatedAt"`       // 1.630488978e+12
			} `json:"Releases"`
			TotalCount int64 `json:"TotalCount"` // 1
		} `json:"ReleasePageList"`
		RequestID string `json:"RequestId"` // d88b989a-ca52-ed03-3cd6-3fea39c81f2f
	} `json:"Response"`
}

// DescribeGitReleases 查询仓库的版本列表
func (r *ReleaseClient) DescribeGitReleases(req DescribeGitReleasesReq) (resp DescribeGitReleasesResp, err error) {
	err = r.Call(&req, &resp)
	return
}
