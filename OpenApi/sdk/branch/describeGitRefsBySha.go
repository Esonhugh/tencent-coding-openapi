package branch

type DescribeGitRefsByShaReq struct {
	Action  string `json:"Action"`  // DescribeGitRefsBySha
	DepotID int64  `json:"DepotId"` // 2
	Sha     string `json:"Sha"`     // 575489c9aed1698d11f1ed398e3c4cb3d45d2ca9
	Type    string `json:"Type"`    // all
}

func (req *DescribeGitRefsByShaReq) SetAction() string {
	req.Action = "DescribeGitRefsBySha"
	return req.Action
}

type DescribeGitRefsByShaResp struct {
	Response struct {
		Refs []struct {
			Name string `json:"Name"` // master
			Type string `json:"Type"` // branch
		} `json:"Refs"`
		RequestID string `json:"RequestId"` // f2faf7f9-5e3b-8b6a-f10d-0bdc34493ff9
	} `json:"Response"`
}

// DescribeGitRefsBySha 查询含有某次提交的标签或分支列表
func (b *BranchClient) DescribeGitRefsBySha(req DescribeGitRefsByShaReq) (resp DescribeGitRefsByShaResp, err error) {
	err = b.Call(&req, &resp)
	return
}
