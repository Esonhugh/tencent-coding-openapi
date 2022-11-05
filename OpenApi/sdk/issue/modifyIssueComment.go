package issue

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/define"
)

type ModifyIssueCommentReq struct {
	Action      string `json:"Action"`      // CreateIssueComment
	CommentID   int64  `json:"CommentId"`   // 1
	Content     string `json:"Content"`     // 这个要配合测试
	IssueCode   int64  `json:"IssueCode"`   // 1
	ProjectName string `json:"ProjectName"` // TestProjectDemo
}

type ModifyIssueCommentResp define.CommonResponse

func (req *ModifyIssueCommentReq) SetAction() string {
	req.Action = "ModifyIssueComment"
	return req.Action
}

// ModifyIssueComment 修改事项评论
func (i *IssueClient) ModifyIssueComment(req ModifyIssueCommentReq) (resp ModifyIssueCommentResp, err error) {
	err = i.Call(&req, &resp)
	return
}
