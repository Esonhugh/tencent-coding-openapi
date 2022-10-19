package commit

type CreateGitFilesReq struct {
	Action   string `json:"Action"`  // CreateGitFiles
	DepotID  int64  `json:"DepotId"` // 809883
	GitFiles []struct {
		Content string `json:"Content"` // line1
		Path    string `json:"Path"`    // src/main/hello.java
	} `json:"GitFiles"`
	LastCommitSha string `json:"LastCommitSha"` // cd3e7461b4061346a4803d1494623fc7ffa7f96a
	Message       string `json:"Message"`       // Create a new file named hello111.java
	NewRef        string `json:"NewRef"`        //
	Ref           string `json:"Ref"`           // master
}

func (req *CreateGitFilesReq) SetAction() string {
	req.Action = "CreateGitFiles"
	return req.Action
}

type CreateGitFilesResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // b974da85-10bf-1411-0322-f84ccde0008c
	} `json:"Response"`
}

// CreateGitFiles 新建文件并提交
func (c *CommitClient) CreateGitFiles(req CreateGitFilesReq) (resp CreateGitFilesResp, err error) {
	err = c.Call(&req, &resp)
	return
}
