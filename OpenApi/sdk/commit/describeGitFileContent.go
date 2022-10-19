package commit

type DescribeGitFileContentReq struct {
	Action    string `json:"Action"`    // DescribeGitFileContent
	CommitSha string `json:"CommitSha"` // 1709525b48d070c131da71b0c6c5f6b8cb69501e
	DepotID   int64  `json:"DepotId"`   // 809883
	Path      string `json:"Path"`      // src/main/main.java
}

func (req *DescribeGitFileContentReq) SetAction() string {
	req.Action = "DescribeGitFileContent"
	return req.Action
}

type DescribeGitFileContentResp struct {
	Response struct {
		GitFileContent struct {
			Content string `json:"Content"` // test test test test line1
			IsLfs   bool   `json:"IsLfs"`   // false
			IsText  bool   `json:"IsText"`  // true
		} `json:"GitFileContent"`
		RequestID string `json:"RequestId"` // 6c1ea96a-0b60-c250-1190-e98376def98e
	} `json:"Response"`
}

// DescribeGitFileContent 查询某次提交某文件的内容
func (c *CommitClient) DescribeGitFileContent(req DescribeGitFileContentReq) (resp DescribeGitFileContentResp, err error) {
	err = c.Call(&req, &resp)
	return
}
