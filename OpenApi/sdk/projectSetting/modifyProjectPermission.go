package projectSetting

type ModifyProjectPermissionReq struct {
	Action     string `json:"Action"`     // ModifyProjectPermission
	ActionFlag bool   `json:"ActionFlag"` // false
	ProjectID  int64  `json:"ProjectId"`  // 2
	RoleID     int64  `json:"RoleId"`     // 120
	UserGK     string `json:"UserGK"`     // TDqyhVeWYN
}

func (req *ModifyProjectPermissionReq) SetAction() string {
	req.Action = "ModifyProjectPermission"
	return req.Action
}

type ModifyProjectPermissionResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// ModifyProjectPermission 配置项目成员权限
func (p *ProjectSettingClient) ModifyProjectPermission(req ModifyProjectPermissionReq) (resp ModifyProjectPermissionResp, err error) {
	err = p.Call(&req, &resp)
	return
}
