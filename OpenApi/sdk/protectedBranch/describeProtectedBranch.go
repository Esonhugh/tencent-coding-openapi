package protectedBranch

type DescribeProtectedBranchReq struct {
	Action     string `json:"Action"`     // DescribeProtectedBranch
	BranchName string `json:"BranchName"` // master
	DepotID    int64  `json:"DepotId"`    // 5001
}

func (req *DescribeProtectedBranchReq) SetAction() string {
	req.Action = "DescribeProtectedBranch"
	return req.Action
}

type DescribeProtectedBranchResp struct {
	Response struct {
		ProtectedBranch struct {
			CommitDate    int64  `json:"CommitDate"`    // 1.60127938e+12
			DenyForcePush bool   `json:"DenyForcePush"` // true
			ForceSquash   bool   `json:"ForceSquash"`   // false
			Name          string `json:"Name"`          // master
			StatusCheck   bool   `json:"StatusCheck"`   // false
		} `json:"ProtectedBranch"`
		RequestID string `json:"RequestId"` // 129c2dc2-223e-4ee6-0643-26a50fb4b5ba
	} `json:"Response"`
}

// DescribeProtectedBranch 查询保护分支状态
func (p *ProtectedBranchClient) DescribeProtectedBranch(req DescribeProtectedBranchReq) (resp DescribeProtectedBranchResp, err error) {
	err = p.Call(&req, &resp)
	return
}
