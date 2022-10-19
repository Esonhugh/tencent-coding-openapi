package git

type DeleteDepotFilePushRuleDenyPrivilegeReq struct {
	Action         string `json:"Action"`         // DeleteDepotFilePushRuleDenyPrivilege
	DepotPath      string `json:"DepotPath"`      // codingcorp/project-d/depot-1
	FilePushRuleID int64  `json:"FilePushRuleId"` // 5
	IsRole         bool   `json:"IsRole"`         // false
	IsUser         bool   `json:"IsUser"`         // true
	UserGlobalKey  string `json:"UserGlobalKey"`  // coding-coding
}

func (req *DeleteDepotFilePushRuleDenyPrivilegeReq) SetAction() string {
	req.Action = "DeleteDepotFilePushRuleDenyPrivilege"
	return req.Action
}

type DeleteDepotFilePushRuleDenyPrivilegeResp struct {
	Action         string `json:"Action"`         // DeleteDepotFilePushRuleDenyPrivilege
	DepotPath      string `json:"DepotPath"`      // codingcorp/project-d/depot-1
	FilePushRuleID int64  `json:"FilePushRuleId"` // 5
	IsRole         bool   `json:"IsRole"`         // true
	IsUser         bool   `json:"IsUser"`         // false
	RoleID         int64  `json:"RoleId"`         // 4
}

// DeleteDepotFilePushRuleDenyPrivilege 删除 git 仓库特权者文件推送权限
func (g *GitClient) DeleteDepotFilePushRuleDenyPrivilege(req DeleteDepotFilePushRuleDenyPrivilegeReq) (resp DeleteDepotFilePushRuleDenyPrivilegeResp, err error) {
	err = g.Call(&req, &resp)
	return
}
