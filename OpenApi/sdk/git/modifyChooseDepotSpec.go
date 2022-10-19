package git

type ModifyChooseDepotSpecReq struct {
	Action        string `json:"Action"`        // ModifyChooseDepotSpec
	DepotPath     string `json:"DepotPath"`     // team-name/project-name/depot-name
	DepotSpecName string `json:"DepotSpecName"` // Gitflow 分支规范
}

func (req *ModifyChooseDepotSpecReq) SetAction() string {
	req.Action = "ModifyChooseDepotSpec"
	return req.Action
}

type ModifyChooseDepotSpecResp struct {
	Response struct {
		Detail struct {
			AllowPushWildRef    bool `json:"AllowPushWildRef"` // false
			DepotBranchTypeList []struct {
				Description string `json:"Description"` // 与仓库设置 > 分支设置中的默认分支保持一致。
				Name        string `json:"Name"`        // 主干分支
				Rule        string `json:"Rule"`        //
			} `json:"DepotBranchTypeList"`
			DepotMergeRequestRuleList []struct {
				SourceBranchTypeName string `json:"SourceBranchTypeName"` // 发布分支
				SourceBranchTypeRule string `json:"SourceBranchTypeRule"` // release/*
				TargetBranchTypeName string `json:"TargetBranchTypeName"` // 主干分支
				TargetBranchTypeRule string `json:"TargetBranchTypeRule"` //
			} `json:"DepotMergeRequestRuleList"`
			Description         string `json:"Description"`         // 支持 master、develop、feature/*、 release/*、hotfix/* 分支
			Name                string `json:"Name"`                // Gitflow 分支规范
			Type                string `json:"Type"`                // system
			UseExistingSolution bool   `json:"UseExistingSolution"` // false
		} `json:"Detail"`
		RequestID string `json:"RequestId"` // 5a344837-42e6-4907-877d-feec15db9a56
	} `json:"Response"`
}

// ModifyChooseDepotSpec 使用、取消使用仓库规范
func (g *GitClient) ModifyChooseDepotSpec(req ModifyChooseDepotSpecReq) (resp ModifyChooseDepotSpecResp, err error) {
	err = g.Call(&req, &resp)
	return
}
