package git

type CreateGitCommitReq struct {
	Action      string `json:"Action"` // CreateGitCommit
	CommitFiles []struct {
		Content string `json:"Content"` // Hello

		NewPath    string `json:"NewPath"`    //
		Path       string `json:"Path"`       // pp.js
		WantDelete bool   `json:"WantDelete"` // false
		WantRename bool   `json:"WantRename"` // false
	} `json:"CommitFiles"`
	DepotPath     string `json:"DepotPath"`     // codingcorp/project-d/depot-1
	LastCommitSha string `json:"LastCommitSha"` //
	Message       string `json:"Message"`       // first commit
	NewRef        string `json:"NewRef"`        //
	Ref           string `json:"Ref"`           // HEAD
}

func (req *CreateGitCommitReq) SetAction() string {
	req.Action = "CreateGitCommit"
	return req.Action
}

type CreateGitCommitResp struct {
	Response struct {
		GitCommit struct {
			CommitDate int64 `json:"CommitDate"` // 1.664266958e+12
			Committer  struct {
				Email string `json:"Email"` // coding@coding.com
				Name  string `json:"Name"`  // coding
			} `json:"Committer"`
			Sha          string `json:"Sha"`          // bb15136ff66c718743daa1a0e13c082a4a54e63b
			ShortMessage string `json:"ShortMessage"` // first commit

		} `json:"GitCommit"`
		RequestID string `json:"RequestId"` // 76214af5-11a2-42a1-b95c-c8ff28a10656
	} `json:"Response"`
}

// CreateGitCommit 创建 git 提交
func (g *GitClient) CreateGitCommit(req CreateGitCommitReq) (resp CreateGitCommitResp, err error) {
	err = g.Call(&req, &resp)
	return
}
