package commit

type ModifyGitCommitRevertReq struct {
	Action     string `json:"Action"`     // ModifyGitCommitRevert
	BranchName string `json:"BranchName"` // feat/mr1
	DepotID    int64  `json:"DepotId"`    // 2
	Message    string `json:"Message"`    // revert this commit
	Sha        string `json:"Sha"`        // e06ebced0f890db2265b8362b3e3f24f30374413
}

func (req *ModifyGitCommitRevertReq) SetAction() string {
	req.Action = "ModifyGitCommitRevert"
	return req.Action
}

type ModifyGitCommitRevertResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 62f19681-d01d-0560-6a92-53256470a6ce
	} `json:"Response"`
}

// ModifyGitCommitRevert 还原某次提交
func (c *CommitClient) ModifyGitCommitRevert(req ModifyGitCommitRevertReq) (resp ModifyGitCommitRevertResp, err error) {
	err = c.Call(&req, &resp)
	return
}
