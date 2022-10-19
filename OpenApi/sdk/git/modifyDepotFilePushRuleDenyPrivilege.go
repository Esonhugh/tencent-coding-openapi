package git

type ModifyDepotFilePushRuleDenyPrivilegeReq struct {
	Action         string `json:"Action"`         // ModifyDepotFilePushRuleDenyPrivilege
	DepotPath      string `json:"DepotPath"`      // codingcorp/project-d/depot-1
	FilePushRuleID int64  `json:"FilePushRuleId"` // 5
	IsDeny         bool   `json:"IsDeny"`         // true
	IsRole         bool   `json:"IsRole"`         // false
	IsUser         bool   `json:"IsUser"`         // true
	UserGlobalKey  string `json:"UserGlobalKey"`  // coding-coding
}

func (req *ModifyDepotFilePushRuleDenyPrivilegeReq) SetAction() string {
	req.Action = "ModifyDepotFilePushRuleDenyPrivilege"
	return req.Action
}

type ModifyDepotFilePushRuleDenyPrivilegeResp struct {
	Action         string `json:"Action"`         // ModifyDepotFilePushRuleDenyPrivilege
	DepotPath      string `json:"DepotPath"`      // codingcorp/project-d/depot-1
	FilePushRuleID int64  `json:"FilePushRuleId"` // 5
	IsDeny         bool   `json:"IsDeny"`         // true
	IsRole         bool   `json:"IsRole"`         // true
	IsUser         bool   `json:"IsUser"`         // false
	RoleID         int64  `json:"RoleId"`         // 4
}

// ModifyDepotFilePushRuleDenyPrivilege 修改 git 仓库特权者文件推送权限
func (g *GitClient) ModifyDepotFilePushRuleDenyPrivilege(req ModifyDepotFilePushRuleDenyPrivilegeReq) (resp ModifyDepotFilePushRuleDenyPrivilegeResp, err error) {
	err = g.Call(&req, &resp)
	return
}
