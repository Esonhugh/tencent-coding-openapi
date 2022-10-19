package branch

type ModifyGitCherryPickReq struct {
	Action     string `json:"Action"`     // ModifyGitCherryPick
	BranchName string `json:"BranchName"` // dev
	DepotID    int64  `json:"DepotId"`    // 2
	Sha        string `json:"Sha"`        // f5b6163521c1dc562cf827af419137142e45f682
}

func (req *ModifyGitCherryPickReq) SetAction() string {
	req.Action = "ModifyGitCherryPick"
	return req.Action
}

type ModifyGitCherryPickResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 8b048fe5-eb93-e8aa-a239-b864d9abd660
	} `json:"Response"`
}

// ModifyGitCherryPick 将某次提交cherry-pick到指定分支
func (b *BranchClient) ModifyGitCherryPick(req ModifyGitCherryPickReq) (resp ModifyGitCherryPickResp, err error) {
	err = b.Call(&req, &resp)
	return
}
