package issue

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/define"
)

type CreateIssueCommentReq struct {
	Action      string `json:"Action"`      // CreateIssueComment
	Content     string `json:"Content"`     // 这个要配合测试
	IssueCode   int64  `json:"IssueCode"`   // 1
	ParentID    int64  `json:"ParentId"`    // 0
	ProjectName string `json:"ProjectName"` // TestProjectDemo
}

type CreateIssueCommentResp define.CommonResponse

func (req *CreateIssueCommentReq) SetAction() string {
	req.Action = "CreateIssueComment"
	return req.Action
}

// CreateIssueComment 创建事项评论
func (i *IssueClient) CreateIssueComment(req CreateIssueCommentReq) (resp CreateIssueCommentResp, err error) {
	err = i.Call(&req, &resp)
	return
}
