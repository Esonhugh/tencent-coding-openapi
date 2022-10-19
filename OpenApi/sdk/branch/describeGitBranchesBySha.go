package branch

type DescribeGitBranchesByShaReq struct {
	Action  string `json:"Action"`  // DescribeGitBranchesBySha
	DepotID int64  `json:"DepotId"` // 1001
	Sha     string `json:"Sha"`     // bf778a27fa30a889a30af6362ba1f16a48dd58dd
}

func (req *DescribeGitBranchesByShaReq) SetAction() string {
	req.Action = "DescribeGitBranchesBySha"
	return req.Action
}

type DescribeGitBranchesByShaResp struct {
	Response struct {
		Refs []struct {
			Ref string `json:"Ref"` // master
		} `json:"Refs"`
		RequestID string `json:"RequestId"` // f16f5a85-9b8d-42f3-84aa-d9061941bad5
	} `json:"Response"`
}

// DescribeGitBranchesBySha 查询提交所在分支
func (b *BranchClient) DescribeGitBranchesBySha(req DescribeGitBranchesByShaReq) (resp DescribeGitBranchesByShaResp, err error) {
	err = b.Call(&req, &resp)
	return
}
