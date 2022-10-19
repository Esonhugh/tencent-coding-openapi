package tag

type DeleteGitProtectedTagRuleReq struct {
	Action  string `json:"Action"`  // DeleteGitProtectedTagRule
	DepotID int64  `json:"DepotId"` // 2
	Rule    string `json:"Rule"`    // 2021*
}

func (req *DeleteGitProtectedTagRuleReq) SetAction() string {
	req.Action = "DeleteGitProtectedTagRule"
	return req.Action
}

type DeleteGitProtectedTagRuleResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // caf50fa3-0c4a-c22f-5932-8d4bf0972c75
	} `json:"Response"`
}

// DeleteGitProtectedTagRule 删除标签保护规则
func (t *TagClient) DeleteGitProtectedTagRule(req DeleteGitProtectedTagRuleReq) (resp DeleteGitProtectedTagRuleResp, err error) {
	err = t.Call(&req, &resp)
	return
}
