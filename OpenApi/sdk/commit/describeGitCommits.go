package commit

type DescribeGitCommitsReq struct {
	Action     string `json:"Action"`     // DescribeGitCommits
	DepotID    int64  `json:"DepotId"`    // 1001
	EndDate    string `json:"EndDate"`    //
	PageNumber int64  `json:"PageNumber"` // 1
	PageSize   int64  `json:"PageSize"`   // 20
	Path       string `json:"Path"`       //
	Ref        string `json:"Ref"`        // master
	StartDate  string `json:"StartDate"`  //
}

func (req *DescribeGitCommitsReq) SetAction() string {
	req.Action = "DescribeGitCommits"
	return req.Action
}

type DescribeGitCommitsResp struct {
	Response struct {
		Commits []struct {
			CommitDate int64 `json:"CommitDate"` // 1.602817039e+12
			Commiter   struct {
				Email string `json:"Email"` // demo1@coding.com
				Name  string `json:"Name"`  // 洋葱猴
			} `json:"Commiter"`
			Sha          string `json:"Sha"`          // bf778a27fa30a889a30af6362ba1f16a48dd58dd
			ShortMessage string `json:"ShortMessage"` // commit 1
		} `json:"Commits"`
		RequestID string `json:"RequestId"` // ae27609d-83a4-6a31-788c-c8d5c2ea9278
	} `json:"Response"`
}

// DescribeGitCommits 查询仓库分支下提交列表
func (c *CommitClient) DescribeGitCommits(req DescribeGitCommitsReq) (resp DescribeGitCommitsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
