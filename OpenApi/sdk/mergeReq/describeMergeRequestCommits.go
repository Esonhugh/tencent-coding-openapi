package mergeReq

type DescribeMergeRequestCommitsReq struct {
	Action  string `json:"Action"`  // DescribeMergeRequestCommits
	DepotID int64  `json:"DepotId"` // 46
	MergeID int64  `json:"MergeId"` // 5
}

func (req *DescribeMergeRequestCommitsReq) SetAction() string {
	req.Action = "DescribeMergeRequestCommits"
	return req.Action
}

type DescribeMergeRequestCommitsResp struct {
	Response struct {
		Commits []struct {
			CommitDate int64 `json:"CommitDate"` // 1.604478235e+12
			Commiter   struct {
				Email string `json:"Email"` // coding@coding.com
				Name  string `json:"Name"`  // 洋葱猴1
			} `json:"Commiter"`
			Sha          string `json:"Sha"`          // 7c86071af40405abad16c90dcf6cb72abcf8184d
			ShortMessage string `json:"ShortMessage"` // commit-1

		} `json:"Commits"`
		RequestID string `json:"RequestId"` // e846cfdb-e842-08a5-7fa3-a484d63fa250
	} `json:"Response"`
}

// DescribeMergeRequestCommits 获取合并请求提交记录
func (m *MergeReqClient) DescribeMergeRequestCommits(req DescribeMergeRequestCommitsReq) (resp DescribeMergeRequestCommitsResp, err error) {
	err = m.Call(&req, &resp)
	return
}
