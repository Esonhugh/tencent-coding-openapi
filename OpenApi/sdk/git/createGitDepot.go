package git

type CreateGitDepotReq struct {
	Action      string `json:"Action"`      // CreateGitDepot
	DepotName   string `json:"DepotName"`   // my-demo
	Description string `json:"Description"` // 啦啦啦啦啦
	ProjectID   int64  `json:"ProjectId"`   // 5001
	Shared      bool   `json:"Shared"`      // true
}

func (req *CreateGitDepotReq) SetAction() string {
	req.Action = "CreateGitDepot"
	return req.Action
}

type CreateGitDepotResp struct {
	Response struct {
		DepotID   int64  `json:"DepotId"`   // 1001
		RequestID string `json:"RequestId"` // 4cc3caaa-ed3d-e610-d274-63694db9e81d
	} `json:"Response"`
}

// CreateGitDepot 创建代码仓库
func (g *GitClient) CreateGitDepot(req CreateGitDepotReq) (resp CreateGitDepotResp, err error) {
	err = g.Call(&req, &resp)
	return
}
