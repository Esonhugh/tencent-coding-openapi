package projectSetting

type DescribeProjectRolesReq struct {
	Action    string `json:"Action"`           // DescribeProjectRoles
	ProjectID int64  `json:"ProjectId,string"` // 1
}

func (req *DescribeProjectRolesReq) SetAction() string {
	req.Action = "DescribeProjectRoles"
	return req.Action
}

type DescribeProjectRolesResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // de0438d0-1e5e-098f-872a-c96863ed3510
		Roles     []struct {
			RoleID       int64  `json:"RoleId"`       // 1
			RoleType     string `json:"RoleType"`     // ProjectMember
			RoleTypeName string `json:"RoleTypeName"` // 开发
		} `json:"Roles"`
	} `json:"Response"`
}

// DescribeProjectRoles 查询项目用户组
func (p *ProjectSettingClient) DescribeProjectRoles(req DescribeProjectRolesReq) (resp DescribeProjectRolesResp, err error) {
	err = p.Call(&req, &resp)
	return
}
