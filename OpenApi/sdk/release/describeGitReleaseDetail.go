package release

type DescribeGitReleaseDetailReq struct {
	Action  string `json:"Action"`  // DescribeGitReleaseDetail
	DepotID int64  `json:"DepotId"` // 809883
	Iid     int64  `json:"Iid"`     // 17
}

func (req *DescribeGitReleaseDetailReq) SetAction() string {
	req.Action = "DescribeGitReleaseDetail"
	return req.Action
}

type DescribeGitReleaseDetailResp struct {
	Response struct {
		Release struct {
			Body            string `json:"Body"`            // Description
			CommitSha       string `json:"CommitSha"`       // 6efc794f0d694fed6cd033ddc984026887af258a
			CreatedAt       int64  `json:"CreatedAt"`       // 1.630487449e+12
			CreatorID       int64  `json:"CreatorId"`       // 1
			DepotID         int64  `json:"DepotId"`         // 809883
			Html            string `json:"Html"`            // <p>Description</p>
			ID              int64  `json:"Id"`              // 31585
			Iid             int64  `json:"Iid"`             // 17
			Pre             bool   `json:"Pre"`             // true
			ProjectID       int64  `json:"ProjectId"`       // 797199
			TagName         string `json:"TagName"`         // TagName
			TargetCommitish string `json:"TargetCommitish"` // 6efc794f0d694fed6cd033ddc984026887af258a
			Title           string `json:"Title"`           // Title
			UpdatedAt       int64  `json:"UpdatedAt"`       // 1.630487449e+12
		} `json:"Release"`
		RequestID string `json:"RequestId"` // 1bfcce2c-6a4b-6e47-5012-cd33ccf2fbe2
	} `json:"Response"`
}

// DescribeGitReleaseDetail 查询仓库的版本详情
func (r *ReleaseClient) DescribeGitReleaseDetail(req DescribeGitReleaseDetailReq) (resp DescribeGitReleaseDetailResp, err error) {
	err = r.Call(&req, &resp)
	return
}
