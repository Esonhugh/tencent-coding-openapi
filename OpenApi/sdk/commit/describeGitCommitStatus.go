package commit

type DescribeGitCommitStatusReq struct {
	Action    string `json:"Action"`    // DescribeGitCommitStatus
	CommitSha string `json:"CommitSha"` // da3e8c8f93516bc8e1609e5a91df98439c6c7b49
	DepotID   int64  `json:"DepotId"`   // 0
	DepotPath string `json:"DepotPath"` // codingcorp/project-d/depot-d
}

func (req *DescribeGitCommitStatusReq) SetAction() string {
	req.Action = "DescribeGitCommitStatus"
	return req.Action
}

type DescribeGitCommitStatusResp struct {
	Response struct {
		RequestID          string `json:"RequestId"` // 20b379d2-5ad8-4b70-ac63-e6cfbc4b354a
		StatusCheckResults []struct {
			CommitSha   string `json:"CommitSha"`   // da3e8c8f93516bc8e1609e5a91df98439c6c7b49
			Context     string `json:"Context"`     // 示例任务 #2
			CreateDate  int64  `json:"CreateDate"`  // 1.663224375e+12
			DepotID     int64  `json:"DepotId"`     // 1
			Description string `json:"Description"` // 构建失败
			State       string `json:"State"`       // COMMIT_STATE_ERROR
			TargetURL   string `json:"TargetURL"`   // http://test.coding.net/p/project-name/ci/job/20684/build/1/pipeline
		} `json:"StatusCheckResults"`
	} `json:"Response"`
}

// DescribeGitCommitStatus 查询提交对应的流水线状态
func (c *CommitClient) DescribeGitCommitStatus(req DescribeGitCommitStatusReq) (resp DescribeGitCommitStatusResp, err error) {
	err = c.Call(&req, &resp)
	return
}
