package commit

type DescribeGitCommitReq struct {
	Action  string `json:"Action"`  // DescribeGitCommit
	DepotID int64  `json:"DepotId"` // 1001
	Sha     string `json:"Sha"`     // bf778a27fa30a889a30af6362ba1f16a48dd58dd
}

func (req *DescribeGitCommitReq) SetAction() string {
	req.Action = "DescribeGitCommit"
	return req.Action
}

type DescribeGitCommitResp struct {
	Response struct {
		GitCommit struct {
			CommitDate int64 `json:"CommitDate"` // 1.602817039e+12
			Commiter   struct {
				Email string `json:"Email"` // mydemo@coding.com
				Name  string `json:"Name"`  // 洋葱猴
			} `json:"Commiter"`
			Sha          string `json:"Sha"`          // bf778a27fa30a889a30af6362ba1f16a48dd58dd
			ShortMessage string `json:"ShortMessage"` // my-commit
		} `json:"GitCommit"`
		RequestID string `json:"RequestId"` // b4fe12f2-c21c-bdd7-5efd-e07dc7f28a31
	} `json:"Response"`
}

// DescribeGitCommit 查询提交详情
func (c *CommitClient) DescribeGitCommit(req DescribeGitCommitReq) (resp DescribeGitCommitResp, err error) {
	err = c.Call(&req, &resp)
	return
}
