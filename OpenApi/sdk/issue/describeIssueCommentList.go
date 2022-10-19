package issue

type DescribeIssueCommentListReq struct {
	Action      string `json:"Action"`      // DescribeIssueCommentList
	IssueCode   int64  `json:"IssueCode"`   // 1
	ProjectName string `json:"ProjectName"` // DemoProjectName
}

func (req *DescribeIssueCommentListReq) SetAction() string {
	req.Action = "DescribeIssueCommentList"
	return req.Action
}

type DescribeIssueCommentListResp struct {
	CommentList []struct {
		CommentID  int64  `json:"CommentId"`  // 1
		Content    string `json:"Content"`    // <a src=http://codingtest.codingcorp.net/p/cc/issues/12> content </a>
		CreatedAt  int64  `json:"CreatedAt"`  // 1.626402309235e+12
		CreatorID  int64  `json:"CreatorId"`  // 1
		ParentID   int64  `json:"ParentId"`   // 0
		RawContent string `json:"RawContent"` // content
		UpdatedAt  int64  `json:"UpdatedAt"`  // 1.626402309235e+12
	} `json:"CommentList"`
}

// DescribeIssueCommentList  查询事项评论列表
func (i *IssueClient) DescribeIssueCommentList(req DescribeIssueCommentListReq) (resp DescribeIssueCommentListResp, err error) {
	err = i.Call(&req, &resp)
	return
}
