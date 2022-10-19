package branch

type DeleteGitMergedBranchesReq struct {
	Action  string `json:"Action"`  // DeleteGitMergedBranches
	DepotID int64  `json:"DepotId"` // 2
}

func (req *DeleteGitMergedBranchesReq) SetAction() string {
	req.Action = "DeleteGitMergedBranches"
	return req.Action
}

type DeleteGitMergedBranchesResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 8b048fe5-eb93-e8aa-a239-b864d9abd660
	} `json:"Response"`
}

// DeleteGitMergedBranches 删除已合并到默认分支的分支（此操作不会删除受保护的分支）
func (b *BranchClient) DeleteGitMergedBranches(req DeleteGitMergedBranchesReq) (resp DeleteGitMergedBranchesResp, err error) {
	err = b.Call(&req, &resp)
	return
}
