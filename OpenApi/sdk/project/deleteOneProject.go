package project

type DeleteOneProjectReq struct {
	Action    string `json:"Action"`    // DeleteOneProject
	ProjectID int64  `json:"ProjectId"` // 2
}

func (req *DeleteOneProjectReq) SetAction() string {
	req.Action = "DeleteOneProject"
	return req.Action
}

type DeleteOneProjectResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// DeleteOneProject 删除项目
func (p *ProjectClient) DeleteOneProject(req DeleteOneProjectReq) (resp DeleteOneProjectResp, err error) {
	err = p.Call(&req, &resp)
	return
}
