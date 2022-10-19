package git

type ModifyDepotLevelDepotSpecReq struct {
	Action    string `json:"Action"`    // ModifyDepotLevelDepotSpec
	DepotPath string `json:"DepotPath"` // codingcorp/project-d/depot-2
	Param     struct {
		AllowPushWildRef    bool `json:"AllowPushWildRef"` // true
		DepotBranchTypeList []struct {
			Description string `json:"Description"` // 这是一个主干分支
			Name        string `json:"Name"`        // 主干分支
			Rule        string `json:"Rule"`        // master
		} `json:"DepotBranchTypeList"`
		DepotMergeRequestRuleList []struct {
			SourceBranchTypeName string `json:"SourceBranchTypeName"` // 版本分支
			TargetBranchTypeName string `json:"TargetBranchTypeName"` // 开发分支
		} `json:"DepotMergeRequestRuleList"`
	} `json:"Param"`
}

func (req *ModifyDepotLevelDepotSpecReq) SetAction() string {
	req.Action = "ModifyDepotLevelDepotSpec"
	return req.Action
}

type ModifyDepotLevelDepotSpecResp struct {
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

// ModifyDepotLevelDepotSpec 修改、新增仓库级别的仓库规范
func (g *GitClient) ModifyDepotLevelDepotSpec(req ModifyDepotLevelDepotSpecReq) (resp ModifyDepotLevelDepotSpecResp, err error) {
	err = g.Call(&req, &resp)
	return
}
