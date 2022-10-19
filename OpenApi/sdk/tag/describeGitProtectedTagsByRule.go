package tag

type DescribeGitProtectedTagsByRuleReq struct {
	Action  string `json:"Action"`  // DescribeGitProtectedTagsByRule
	DepotID int64  `json:"DepotId"` // 2
	Rule    string `json:"Rule"`    // 2022*
}

func (req *DescribeGitProtectedTagsByRuleReq) SetAction() string {
	req.Action = "DescribeGitProtectedTagsByRule"
	return req.Action
}

type DescribeGitProtectedTagsByRuleResp struct {
	Response struct {
		GitTags []struct {
			Message string `json:"Message"` // delete file test/test.dart
			TagName string `json:"TagName"` // 2022.01.12
		} `json:"GitTags"`
		RequestID string `json:"RequestId"` // 3d9b3771-9b56-7da8-a31d-9e5bb402d659
	} `json:"Response"`
}

// DescribeGitProtectedTagsByRule 根据标签保护规则查询受保护的标签列表
func (t *TagClient) DescribeGitProtectedTagsByRule(req DescribeGitProtectedTagsByRuleReq) (resp DescribeGitProtectedTagsByRuleResp, err error) {
	err = t.Call(&req, &resp)
	return
}
