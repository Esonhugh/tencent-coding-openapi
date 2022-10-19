package branchProtection

type ModifyBranchProtectionReq struct {
	Action           string `json:"Action"` // ModifyBranchProtection
	BranchProtection struct {
		BranchProtectionID     int64  `json:"BranchProtectionId"`     // 28250
		DenyForcePush          bool   `json:"DenyForcePush"`          // true
		ForceSquash            bool   `json:"ForceSquash"`            // false
		MatcherCount           int64  `json:"MatcherCount"`           // 0
		RequiredCodeOwnerGrant bool   `json:"RequiredCodeOwnerGrant"` // false
		RequiredGrantCount     int64  `json:"RequiredGrantCount"`     // 1
		RequiredStatusCheck    bool   `json:"RequiredStatusCheck"`    // false
		Rule                   string `json:"Rule"`                   // dev/*
		SrcMustMergedTo        string `json:"SrcMustMergedTo"`        //
	} `json:"BranchProtection"`
	DepotID int64 `json:"DepotId"` // 809883
}

func (req *ModifyBranchProtectionReq) SetAction() string {
	req.Action = "ModifyBranchProtection"
	return req.Action
}

type ModifyBranchProtectionResp struct {
	Response struct {
		BranchProtection struct {
			BranchProtectionID     int64  `json:"BranchProtectionId"`     // 28250
			DenyForcePush          bool   `json:"DenyForcePush"`          // true
			ForceSquash            bool   `json:"ForceSquash"`            // false
			MatcherCount           int64  `json:"MatcherCount"`           // 0
			RequiredCodeOwnerGrant bool   `json:"RequiredCodeOwnerGrant"` // false
			RequiredGrantCount     int64  `json:"RequiredGrantCount"`     // 1
			RequiredStatusCheck    bool   `json:"RequiredStatusCheck"`    // false
			Rule                   string `json:"Rule"`                   // dev/*
			SrcMustMergedTo        string `json:"SrcMustMergedTo"`        //
		} `json:"BranchProtection"`
		RequestID string `json:"RequestId"` // 93bd76b6-68aa-af54-2607-b637e53edb74
	} `json:"Response"`
}

// ModifyBranchProtection 修改保护分支规则
func (b *BranchProtectionClient) ModifyBranchProtection(req ModifyBranchProtectionReq) (resp ModifyBranchProtectionResp, err error) {
	err = b.Call(&req, &resp)
	return
}
