package git

type DescribeGitFileReq struct {
	Action  string `json:"Action"`  // DescribeGitFile
	DepotID int64  `json:"DepotId"` // 1
	Path    string `json:"Path"`    // aaa
	Ref     string `json:"Ref"`     // master
}

func (req *DescribeGitFileReq) SetAction() string {
	req.Action = "DescribeGitFile"
	return req.Action
}

type DescribeGitFileResp struct {
	Response struct {
		GitFile struct {
			Content       string `json:"Content"`       // YmJiCgrlpJrlkIPngrnlkIPnmoToj5w=
			ContentSha256 string `json:"ContentSha256"` // 3073271fa083e505836c873ffd90be5e7673c7d4787174dd7157ad360b58bbdc
			Encoding      string `json:"Encoding"`      // base64
			FileName      string `json:"FileName"`      // bbb
			FilePath      string `json:"FilePath"`      // src/bbb
			Sha           string `json:"Sha"`           // 45d1b22600241eda05736401950d65cfd23efa13
			Size          int64  `json:"Size"`          // 23
		} `json:"GitFile"`
		RequestID string `json:"RequestId"` // 56c5f669-93c7-8088-9d2f-a8872fc879ed
	} `json:"Response"`
}

// DescribeGitFile 获取文件详情
func (g *GitClient) DescribeGitFile(req DescribeGitFileReq) (resp DescribeGitFileResp, err error) {
	err = g.Call(&req, &resp)
	return
}
