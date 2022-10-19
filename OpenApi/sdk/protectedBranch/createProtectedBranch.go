package protectedBranch

type CreateProtectedBranchReq struct {
	Action     string `json:"Action"`     // CreateProtectedBranch
	BranchName string `json:"BranchName"` // master
	DepotID    int64  `json:"DepotId"`    // 5001
}

func (req *CreateProtectedBranchReq) SetAction() string {
	req.Action = "CreateProtectedBranch"
	return req.Action
}

type CreateProtectedBranchResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 2cf6e3b6-7fa7-0433-50b1-73e6226ac15b
	} `json:"Response"`
}

// CreateProtectedBranch 设置保护分支
func (p *ProtectedBranchClient) CreateProtectedBranch(req CreateProtectedBranchReq) (resp CreateProtectedBranchResp, err error) {
	err = p.Call(&req, &resp)
	return
}
