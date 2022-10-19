package branchProtection

type CreateBranchProtectionMemberReq struct {
	Action             string `json:"Action"`             // CreateBranchProtectionMember
	AllowPush          bool   `json:"AllowPush"`          // true
	BranchProtectionID int64  `json:"BranchProtectionId"` // 28234
	DepotID            int64  `json:"DepotId"`            // 809883
	UserGlobalKey      string `json:"UserGlobalKey"`      // coding
}

func (req *CreateBranchProtectionMemberReq) SetAction() string {
	req.Action = "CreateBranchProtectionMember"
	return req.Action
}

type CreateBranchProtectionMemberResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 51e123c8-48af-72f9-ac98-dc32f393355f
	} `json:"Response"`
}

// CreateBranchProtectionMember 新建保护分支规则的成员
func (b *BranchProtectionClient) CreateBranchProtectionMember(req CreateBranchProtectionMemberReq) (resp CreateBranchProtectionMemberResp, err error) {
	err = b.Call(&req, &resp)
	return
}
