package tag

type CreateGitProtectedTagRuleReq struct {
	Action  string `json:"Action"`  // CreateGitProtectedTagRule
	DepotID int64  `json:"DepotId"` // 2
	Rule    string `json:"Rule"`    // 2022*
}

func (req *CreateGitProtectedTagRuleReq) SetAction() string {
	req.Action = "CreateGitProtectedTagRule"
	return req.Action
}

type CreateGitProtectedTagRuleResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 3d9b3771-9b56-7da8-a31d-9e5bb402d659
	} `json:"Response"`
}

// CreateGitProtectedTagRule 创建标签保护规则
func (t *TagClient) CreateGitProtectedTagRule(req CreateGitProtectedTagRuleReq) (resp CreateGitProtectedTagRuleResp, err error) {
	err = t.Call(&req, &resp)
	return
}
