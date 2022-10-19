package tag

type DescribeGitTagsReq struct {
	Action  string `json:"Action"`  // DescribeGitTags
	DepotID int64  `json:"DepotId"` // 1001
}

func (req *DescribeGitTagsReq) SetAction() string {
	req.Action = "DescribeGitTags"
	return req.Action
}

type DescribeGitTagsResp struct {
	Response struct {
		GitTags []struct {
			Message string `json:"Message"` // this is tag-demo1
			TagName string `json:"TagName"` // tag-demo1
		} `json:"GitTags"`
		RequestID string `json:"RequestId"` // 05d4f18d-8602-ad9c-f6a2-8ad6f8f73127
	} `json:"Response"`
}

// DescribeGitTags 查询仓库所有标签
func (t *TagClient) DescribeGitTags(req DescribeGitTagsReq) (resp DescribeGitTagsResp, err error) {
	err = t.Call(&req, &resp)
	return
}
