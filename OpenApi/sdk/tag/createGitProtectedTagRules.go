package tag

type CreateGitProtectedTagRulesReq struct {
	Action  string   `json:"Action"`  // CreateGitProtectedTagRules
	DepotID int64    `json:"DepotId"` // 2
	Rule    []string `json:"Rule"`    // 2022.01*
}

func (req *CreateGitProtectedTagRulesReq) SetAction() string {
	req.Action = "CreateGitProtectedTagRules"
	return req.Action
}

type CreateGitProtectedTagRulesResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // e7251f89-b65c-258b-b96e-62e25a2ef74c
	} `json:"Response"`
}

// CreateGitProtectedTagRules 批量创建标签保护规则
func (t *TagClient) CreateGitProtectedTagRules(req CreateGitProtectedTagRulesReq) (resp CreateGitProtectedTagRulesResp, err error) {
	err = t.Call(&req, &resp)
	return
}
