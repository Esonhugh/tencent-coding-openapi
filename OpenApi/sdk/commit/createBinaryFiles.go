package commit

type CreateBinaryFilesReq struct {
	Action   string `json:"Action"`  // CreateBinaryFiles
	DepotID  int64  `json:"DepotId"` // 809883
	DestRef  string `json:"DestRef"` // master
	GitFiles []struct {
		Content string `json:"Content"` // SSUyN20lMjBhJTIwYmluYXJ5JTIwZmlsZSUyMHRoYXQlMjBuZWVkcyUyMHRvJTIwYmUlMjBkZWNyeXB0ZWQu
		NewPath string `json:"NewPath"` //
		Path    string `json:"Path"`    // src/main/binary2.test
	} `json:"GitFiles"`
	LastCommitSha string `json:"LastCommitSha"` // 1db4eabbded90a7669094b3e5162d9098840a2cf
	Message       string `json:"Message"`       // create a binary file2
	SrcRef        string `json:"SrcRef"`        // master
	UserID        int64  `json:"UserId"`        // 1
}

func (req *CreateBinaryFilesReq) SetAction() string {
	req.Action = "CreateBinaryFiles"
	return req.Action
}

type CreateBinaryFilesResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 10623f03-0f5a-0616-239f-21711e95eb4e
	} `json:"Response"`
}

// CreateBinaryFiles 新建内容是二进制的文件
func (c *CommitClient) CreateBinaryFiles(req CreateBinaryFilesReq) (resp CreateBinaryFilesResp, err error) {
	err = c.Call(&req, &resp)
	return
}
