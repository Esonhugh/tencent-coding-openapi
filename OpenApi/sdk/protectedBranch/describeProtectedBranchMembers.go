package protectedBranch

type DescribeProtectedBranchMembersReq struct {
	Action  string `json:"Action"`  // DescribeProtectedBranchMembers
	DepotID int64  `json:"DepotId"` // 5001
}

func (req *DescribeProtectedBranchMembersReq) SetAction() string {
	req.Action = "DescribeProtectedBranchMembers"
	return req.Action
}

type DescribeProtectedBranchMembersResp struct {
	Response struct {
		Members []struct {
			GlobalKey       string `json:"GlobalKey"`       // coding-demo
			HasPushAccess   bool   `json:"HasPushAccess"`   // true
			HasUpdateAccess bool   `json:"HasUpdateAccess"` // true
			Name            string `json:"Name"`            // 洋葱猴
		} `json:"Members"`
		RequestID string `json:"RequestId"` // 9aaf9eb9-f9fd-696b-ca60-4a6d4846ca3c
	} `json:"Response"`
}

// DescribeProtectedBranchMembers 查询所有保护分支成员
func (p *ProtectedBranchClient) DescribeProtectedBranchMembers(req DescribeProtectedBranchMembersReq) (resp DescribeProtectedBranchMembersResp, err error) {
	err = p.Call(&req, &resp)
	return
}
