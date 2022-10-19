package branchProtection

type DescribeBranchProtectionReq struct {
	Action             string `json:"Action"`             // DescribeBranchProtection
	BranchProtectionID int64  `json:"BranchProtectionId"` // 28234
	DepotID            int64  `json:"DepotId"`            // 809883
}

func (req *DescribeBranchProtectionReq) SetAction() string {
	req.Action = "DescribeBranchProtection"
	return req.Action
}

type DescribeBranchProtectionResp struct {
	Response struct {
		BranchProtection struct {
			BranchProtectionID     int64  `json:"BranchProtectionId"`     // 28234
			DenyForcePush          bool   `json:"DenyForcePush"`          // true
			ForceSquash            bool   `json:"ForceSquash"`            // false
			MatcherCount           int64  `json:"MatcherCount"`           // 0
			RequiredCodeOwnerGrant bool   `json:"RequiredCodeOwnerGrant"` // false
			RequiredGrantCount     int64  `json:"RequiredGrantCount"`     // 1
			RequiredStatusCheck    bool   `json:"RequiredStatusCheck"`    // false
			Rule                   string `json:"Rule"`                   // dev/*
			SrcMustMergedTo        string `json:"SrcMustMergedTo"`        //
		} `json:"BranchProtection"`
		RequestID string `json:"RequestId"` // c96a77a6-e366-f8b8-9d96-96d413d55414
	} `json:"Response"`
}

// DescribeBranchProtection 查询保护分支规则
func (b *BranchProtectionClient) DescribeBranchProtection(req DescribeBranchProtectionReq) (resp DescribeBranchProtectionResp, err error) {
	err = b.Call(&req, &resp)
	return
}
