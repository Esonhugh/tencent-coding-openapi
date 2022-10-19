package tag

type DescribeGitProtectedTagsReq struct {
	Action  string `json:"Action"`  // DescribeGitProtectedTags
	DepotID int64  `json:"DepotId"` // 2
}

func (req *DescribeGitProtectedTagsReq) SetAction() string {
	req.Action = "DescribeGitProtectedTags"
	return req.Action
}

type DescribeGitProtectedTagsResp struct {
	Response struct {
		GitTags []struct {
			Message string `json:"Message"` //
			TagName string `json:"TagName"` // 2021.12.31
		} `json:"GitTags"`
		RequestID string `json:"RequestId"` // 8b048fe5-eb93-e8aa-a239-b864d9abd660
	} `json:"Response"`
}

// DescribeGitProtectedTags 查询受保护的标签列表
func (t *TagClient) DescribeGitProtectedTags(req DescribeGitProtectedTagsReq) (resp DescribeGitProtectedTagsResp, err error) {
	err = t.Call(&req, &resp)
	return
}
