package git

type CreateDepotByTemplateReq struct {
	Action      string `json:"Action"`      // CreateDepotByTemplate
	DepotName   string `json:"DepotName"`   // hexo-demo-2
	Description string `json:"Description"` // 这是一个测试的 Spring 仓库
	ProjectID   int64  `json:"ProjectId"`   // 450
	Template    string `json:"Template"`    // SPRING
	UserID      int64  `json:"UserId"`      // 1
}

func (req *CreateDepotByTemplateReq) SetAction() string {
	req.Action = "CreateDepotByTemplate"
	return req.Action
}

type CreateDepotByTemplateResp struct {
	Response struct {
		DepotInfo struct {
			HttpsURL  string `json:"HttpsUrl"`  // http://coding.com/codingcorp/demo/hexo-demo-2.git
			ID        int64  `json:"Id"`        // 108
			Name      string `json:"Name"`      // hexo-demo-2
			ProjectID int64  `json:"ProjectId"` // 450
			SshURL    string `json:"SshUrl"`    // git@coding.com:codingcorp/demo/hexo-demo-2.git
			VcsType   string `json:"VcsType"`   // git
			WebURL    string `json:"WebUrl"`    // http://codingcorp.coding.net/p/demo/d/hexo-demo-2
		} `json:"DepotInfo"`
		RequestID string `json:"RequestId"` // c783ce49-5969-7f75-0145-44f694af8f8f
	} `json:"Response"`
}

// CreateDepotByTemplate 根据模板创建仓库
func (g *GitClient) CreateDepotByTemplate(req CreateDepotByTemplateReq) (resp CreateDepotByTemplateResp, err error) {
	err = g.Call(&req, &resp)
	return
}
