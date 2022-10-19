package git

type DeleteGitFilesReq struct {
	Action        string   `json:"Action"`        // DeleteGitFiles
	CommitMessage string   `json:"CommitMessage"` // delete file test/test.dart
	DepotID       int64    `json:"DepotId"`       // 2
	Paths         []string `json:"Paths"`         // test/test.dart
	Ref           string   `json:"Ref"`           // master
}

func (req *DeleteGitFilesReq) SetAction() string {
	req.Action = "DeleteGitFiles"
	return req.Action
}

type DeleteGitFilesResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // b81757db-42dc-98f9-0f01-9d7c1abe0f50
	} `json:"Response"`
}

// DeleteGitFiles 在仓库中删除文件
func (g *GitClient) DeleteGitFiles(req DeleteGitFilesReq) (resp DeleteGitFilesResp, err error) {
	err = g.Call(&req, &resp)
	return
}
