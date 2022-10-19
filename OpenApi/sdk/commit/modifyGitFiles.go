package commit

type ModifyGitFilesReq struct {
	Action   string `json:"Action"`  // ModifyGitFiles
	DepotID  int64  `json:"DepotId"` // 809883
	GitFiles []struct {
		Content string `json:"Content"` // line1
		NewPath string `json:"NewPath"` //
		Path    string `json:"Path"`    // src/main/hello.java
	} `json:"GitFiles"`
	LastCommitSha string `json:"LastCommitSha"` // 93c4b9987ba34183ecc2c45eee26aa44dc1c08ea
	Message       string `json:"Message"`       // Modify hello.java
	NewRef        string `json:"NewRef"`        //
	Ref           string `json:"Ref"`           // master
}

func (req *ModifyGitFilesReq) SetAction() string {
	req.Action = "ModifyGitFiles"
	return req.Action
}

type ModifyGitFilesResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // b974da85-10bf-1411-0322-f84ccde0008c
	} `json:"Response"`
}

// ModifyGitFiles 修改文件并提交
func (c *CommitClient) ModifyGitFiles(req ModifyGitFilesReq) (resp ModifyGitFilesResp, err error) {
	err = c.Call(&req, &resp)
	return
}
