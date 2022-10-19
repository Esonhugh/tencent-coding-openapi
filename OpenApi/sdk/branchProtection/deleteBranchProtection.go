package branchProtection

type DeleteBranchProtectionReq struct {
	Action             string `json:"Action"`             // DeleteBranchProtection
	BranchProtectionID int64  `json:"BranchProtectionId"` // 10
	DepotID            int64  `json:"DepotId"`            // 2
}

func (req *DeleteBranchProtectionReq) SetAction() string {
	req.Action = "DeleteBranchProtection"
	return req.Action
}

type DeleteBranchProtectionResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // b81757db-42dc-98f9-0f01-9d7c1abe0f50
	} `json:"Response"`
}

// DeleteBranchProtection 删除保护分支规则
func (b *BranchProtectionClient) DeleteBranchProtection(req DeleteBranchProtectionReq) (resp DeleteBranchProtectionResp, err error) {
	err = b.Call(&req, &resp)
	return
}
