package project

type CreateProjectMemberReq struct {
	Action            string   `json:"Action"`            // CreateProjectMember
	ProjectID         int64    `json:"ProjectId"`         // 1
	Type              int64    `json:"Type"`              // 1
	UserGlobalKeyList []string `json:"UserGlobalKeyList"` // gk1
}

func (req *CreateProjectMemberReq) SetAction() string {
	req.Action = "CreateProjectMember"
	return req.Action
}

type CreateProjectMemberResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// CreateProjectMember 增加项目成员
func (p *ProjectClient) CreateProjectMember(req CreateProjectMemberReq) (resp CreateProjectMemberResp, err error) {
	err = p.Call(&req, &resp)
	return
}
