package branch

type DeleteGitBranchReq struct {
	Action     string `json:"Action"`     // DeleteGitBranch
	BranchName string `json:"BranchName"` // branch-demo
	DepotID    int64  `json:"DepotId"`    // 1001
}

func (req *DeleteGitBranchReq) SetAction() string {
	req.Action = "DeleteGitBranch"
	return req.Action
}

type DeleteGitBranchResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 30892813-210d-a1cf-10bd-ca11cb1e271a
	} `json:"Response"`
}

// DeleteGitBranch 删除分支
func (b *BranchClient) DeleteGitBranch(req DeleteGitBranchReq) (resp DeleteGitBranchResp, err error) {
	err = b.Call(&req, &resp)
	return
}
