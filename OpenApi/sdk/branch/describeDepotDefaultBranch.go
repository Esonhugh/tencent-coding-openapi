package branch

type DescribeDepotDefaultBranchReq struct {
	Action    string `json:"Action"`    // DescribeDepotDefaultBranch
	DepotPath string `json:"DepotPath"` // codingcorp/project-d/depot-d
}

func (req *DescribeDepotDefaultBranchReq) SetAction() string {
	req.Action = "DescribeDepotDefaultBranch"
	return req.Action
}

type DescribeDepotDefaultBranchResp struct {
	Response struct {
		BranchName string `json:"BranchName"` // master
		RequestID  string `json:"RequestId"`  // 20b379d2-5ad8-4b70-ac63-e6cfbc4b354a
	} `json:"Response"`
}

// DescribeDepotDefaultBranch 查询仓库的默认分支
func (b *BranchClient) DescribeDepotDefaultBranch(req DescribeDepotDefaultBranchReq) (resp DescribeDepotDefaultBranchResp, err error) {
	err = b.Call(&req, &resp)
	return
}
