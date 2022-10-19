package branchProtection

type CreateBranchProtectionReq struct {
	Action  string `json:"Action"`  // CreateBranchProtection
	DepotID int64  `json:"DepotId"` // 809883
	Rule    string `json:"Rule"`    // dev/*
}

func (req *CreateBranchProtectionReq) SetAction() string {
	req.Action = "CreateBranchProtection"
	return req.Action
}

type CreateBranchProtectionResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // c20cbf85-844c-6a7d-8f4d-99400db1e9e1
	} `json:"Response"`
}

// CreateBranchProtection 新建保护分支规则
func (b *BranchProtectionClient) CreateBranchProtection(req CreateBranchProtectionReq) (resp CreateBranchProtectionResp, err error) {
	err = b.Call(&req, &resp)
	return
}
