package cd

type BindCdApplicationToProjectReq struct {
	Action      string `json:"Action"`      // BindCdApplicationToProject
	Application string `json:"Application"` // dev
	ProjectName string `json:"ProjectName"` // test
}

func (req *BindCdApplicationToProjectReq) SetAction() string {
	req.Action = "BindCdApplicationToProject"
	return req.Action
}

type BindCdApplicationToProjectResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 501d4e6c-10d9-7b28-326f-0a1ce5abb9e7
	} `json:"Response"`
}

// BindCdApplicationToProject 绑定 CD 应用到项目
func (c *CdClient) BindCdApplicationToProject(req BindCdApplicationToProjectReq) (resp BindCdApplicationToProjectResp, err error) {
	err = c.Call(&req, &resp)
	return
}
