package branch

type DescribeGitBranchesReq struct {
	Action     string `json:"Action"`     // DescribeGitBranches
	DepotID    int64  `json:"DepotId"`    // 1001
	KeyWord    string `json:"KeyWord"`    //
	PageNumber int64  `json:"PageNumber"` // 1
	PageSize   int64  `json:"PageSize"`   // 20
}

func (req *DescribeGitBranchesReq) SetAction() string {
	req.Action = "DescribeGitBranches"
	return req.Action
}

type DescribeGitBranchesResp struct {
	Response struct {
		Branches []struct {
			BranchName      string `json:"BranchName"`      // master
			IsDefaultBranch bool   `json:"IsDefaultBranch"` // true
			IsProtected     bool   `json:"IsProtected"`     // false
			Sha             string `json:"Sha"`             // 92811220f02b7c9f1ef5559c74ced479f271c905
		} `json:"Branches"`
		RequestID  string `json:"RequestId"`  // 7c57d93d-e731-6150-8575-b3bf68f11b05
		TotalCount int64  `json:"TotalCount"` // 4
	} `json:"Response"`
}

// DescribeGitBranches 查询所有分支信息
func (b *BranchClient) DescribeGitBranches(req DescribeGitBranchesReq) (resp DescribeGitBranchesResp, err error) {
	err = b.Call(&req, &resp)
	return
}
