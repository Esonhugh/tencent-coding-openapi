package git

type DescribeGitBlobRawReq struct {
	Action  string `json:"Action"`  // DescribeGitBlobRaw
	BlobSha string `json:"BlobSha"` // a7954ee405100719a59c1e03f4ac98ce92a2a257
	DepotID int64  `json:"DepotId"` // 2
}

func (req *DescribeGitBlobRawReq) SetAction() string {
	req.Action = "DescribeGitBlobRaw"
	return req.Action
}

type DescribeGitBlobRawResp struct {
	Response struct {
		Content   string `json:"Content"`   // name: extended_text
		RequestID string `json:"RequestId"` // 5649c350-f27a-a84a-3cff-f76c36f03280
	} `json:"Response"`
}

// DescribeGitBlobRaw 查询 Git Blob raw 信息
func (g *GitClient) DescribeGitBlobRaw(req DescribeGitBlobRawReq) (resp DescribeGitBlobRawResp, err error) {
	err = g.Call(&req, &resp)
	return
}
