package branch

type DescribeGitMergeBaseReq struct {
	Action  string `json:"Action"`  // DescribeGitMergeBase
	DepotID int64  `json:"DepotId"` // 2
	DestRef string `json:"DestRef"` // dev
	SrcRef  string `json:"SrcRef"`  // master
}

func (req *DescribeGitMergeBaseReq) SetAction() string {
	req.Action = "DescribeGitMergeBase"
	return req.Action
}

type DescribeGitMergeBaseResp struct {
	Response struct {
		Commit struct {
			CommitDate int64 `json:"CommitDate"` // 1.66442925e+12
			Committer  struct {
				Email string `json:"Email"` // coding@coding.com
				Name  string `json:"Name"`  // coding
			} `json:"Committer"`
			Sha          string `json:"Sha"`          // ed8a9bc096be952b983625c1a6770449887d4e78
			ShortMessage string `json:"ShortMessage"` // emo.js

		} `json:"Commit"`
		RequestID string `json:"RequestId"` // 143c9ab0-eccb-436b-9a13-b67c1a415cf4
	} `json:"Response"`
}

// DescribeGitMergeBase 查询两个分支的公共祖先
func (b *BranchClient) DescribeGitMergeBase(req DescribeGitMergeBaseReq) (resp DescribeGitMergeBaseResp, err error) {
	err = b.Call(&req, &resp)
	return
}
