package commit

type DescribeCommitsBetween2CommitReq struct {
	Action  string `json:"Action"`  // DescribeCommitsBetween2Commit
	DepotID int64  `json:"DepotId"` // 5001
	Source  string `json:"Source"`  //
	Target  string `json:"Target"`  //
}

func (req *DescribeCommitsBetween2CommitReq) SetAction() string {
	req.Action = "DescribeCommitsBetween2Commit"
	return req.Action
}

type DescribeCommitsBetween2CommitResp struct {
	Response struct {
		DifferentOfCommitInfo struct {
			Commits []struct {
				CommitDate int64 `json:"CommitDate"` // 1.602668567e+12
				Commiter   struct {
					Email string `json:"Email"` // coding@coding.com
					Name  string `json:"Name"`  // coding
				} `json:"Commiter"`
				Sha          string `json:"Sha"`          // be9a51ea0dfb1a43e1e714db0e0ecad3009bf362
				ShortMessage string `json:"ShortMessage"` // asdasd

			} `json:"Commits"`
			Deletions          int64 `json:"Deletions"` // 0
			DifferentOfCommits []struct {
				Deletions  int64  `json:"Deletions"`  // 0
				Insertions int64  `json:"Insertions"` // 1
				Name       string `json:"Name"`       // dada
				Path       string `json:"Path"`       // dada
			} `json:"DifferentOfCommits"`
			Insertions    int64 `json:"Insertions"`    // 1
			UpdateFileNum int64 `json:"UpdateFileNum"` // 1
		} `json:"DifferentOfCommitInfo"`
		RequestID string `json:"RequestId"` // ae841435-a933-d500-b363-c25d3843bfa3
	} `json:"Response"`
}

// DescribeCommitsBetween2Commit 两次提交之间的提交历史
func (c *CommitClient) DescribeCommitsBetween2Commit(req DescribeCommitsBetween2CommitReq) (resp DescribeCommitsBetween2CommitResp, err error) {
	err = c.Call(&req, &resp)
	return
}
