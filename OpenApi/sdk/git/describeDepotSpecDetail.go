package git

type DescribeDepotSpecDetailReq struct {
	Action        string `json:"Action"`        // DescribeDepotSpecDetail
	DepotPath     string `json:"DepotPath"`     // team-name/project-name/depot-name
	DepotSpecName string `json:"DepotSpecName"` // GitLab 分支规范
}

func (req *DescribeDepotSpecDetailReq) SetAction() string {
	req.Action = "DescribeDepotSpecDetail"
	return req.Action
}

type DescribeDepotSpecDetailResp struct {
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
		RequestID string `json:"RequestId"` // a0b22a96-62b6-4859-8a86-fa88445fac74
	} `json:"Response"`
}

// DescribeDepotSpecDetail 查询仓库规范详情
func (g *GitClient) DescribeDepotSpecDetail(req DescribeDepotSpecDetailReq) (resp DescribeDepotSpecDetailResp, err error) {
	err = g.Call(&req, &resp)
	return
}
