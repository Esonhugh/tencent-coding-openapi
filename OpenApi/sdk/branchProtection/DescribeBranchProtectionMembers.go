package branchProtection

type DescribeBranchProtectionMembersReq struct {
	Action             string `json:"Action"`             // DescribeBranchProtectionMembers
	BranchProtectionID int64  `json:"BranchProtectionId"` // 28250
	DepotID            int64  `json:"DepotId"`            // 809883
}

func (req *DescribeBranchProtectionMembersReq) SetAction() string {
	req.Action = "DescribeBranchProtectionMembers"
	return req.Action
}

type DescribeBranchProtectionMembersResp struct {
	Response struct {
		Members []struct {
			AllowPush bool   `json:"AllowPush"` // true
			GlobalKey string `json:"GlobalKey"` // coding
			Name      string `json:"Name"`      // Coding管理员1
		} `json:"Members"`
		RequestID string `json:"RequestId"` // 3a8c60bd-14ff-65e2-2458-e0f7e92d56b4
	} `json:"Response"`
}

// DescribeBranchProtectionMembers 查询保护分支规则的管理员
func (b *BranchProtectionClient) DescribeBranchProtectionMembers(req DescribeBranchProtectionMembersReq) (resp DescribeBranchProtectionMembersResp, err error) {
	err = b.Call(&req, &resp)
	return
}
