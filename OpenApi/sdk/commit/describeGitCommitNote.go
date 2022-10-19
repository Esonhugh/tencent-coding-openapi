package commit

type DescribeGitCommitNoteReq struct {
	Action    string `json:"Action"`    // DescribeGitCommitNote
	CommitSha string `json:"CommitSha"` // 5178d68bd3ee2a93733354884f0888584d671d7b
	DepotID   int64  `json:"DepotId"`   // 1383
	NotesRef  string `json:"NotesRef"`  //
}

func (req *DescribeGitCommitNoteReq) SetAction() string {
	req.Action = "DescribeGitCommitNote"
	return req.Action
}

type DescribeGitCommitNoteResp struct {
	Response struct {
		CommitNote string `json:"CommitNote"` // coding
		RequestID  string `json:"RequestId"`  // f2212265-890c-fb76-2154-794c2a395610
	} `json:"Response"`
}

// DescribeGitCommitNote 获取提交注释
func (c *CommitClient) DescribeGitCommitNote(req DescribeGitCommitNoteReq) (resp DescribeGitCommitNoteResp, err error) {
	err = c.Call(&req, &resp)
	return
}
