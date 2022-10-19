package issue

type ModifyIssueParentRequirementReq struct {
	IssueCode       int64  `json:"IssueCode"`       // 10
	ParentIssueCode int64  `json:"ParentIssueCode"` // 1
	ProjectName     string `json:"ProjectName"`     // project-demo
}

func (req *ModifyIssueParentRequirementReq) SetAction() string {
	return "ModifyIssueParentRequirement"
}

type ModifyIssueParentRequirementResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ModifyIssueParentRequirement 修改事项父需求
func (i *IssueClient) ModifyIssueParentRequirement(req ModifyIssueParentRequirementReq) (resp ModifyIssueParentRequirementResp, err error) {
	err = i.Call(&req, &resp)
	return
}
