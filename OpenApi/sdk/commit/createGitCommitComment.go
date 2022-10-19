package commit

type CreateGitCommitCommentReq struct {
	Action  string `json:"Action"`       // CreateGitCommitComment
	Content string `json:"Content"`      // Hello
	DepotID int64  `json:"DepotId"`      // 2
	Index   int64  `json:"Index,string"` // 1
	Path    string `json:"Path"`         // lib/src/extended_render_paragraph.dart
	Sha     string `json:"sha"`          // 575489c9aed1698d11f1ed398e3c4cb3d45d2ca9
}

func (req *CreateGitCommitCommentReq) SetAction() string {
	req.Action = "CreateGitCommitComment"
	return req.Action
}

type CreateGitCommitCommentResp struct {
	Response struct {
		Comment struct {
			Author struct {
				Avatar    string `json:"Avatar"`    // /static/fruit_avatar/Fruit-20.png
				Email     string `json:"Email"`     //
				GlobalKey string `json:"GlobalKey"` // coding-coding
				ID        int64  `json:"Id"`        // 2
				Name      string `json:"Name"`      // coding
				Status    string `json:"Status"`    //
				TeamID    int64  `json:"TeamId"`    // 1
			} `json:"Author"`
			CommitSha string `json:"CommitSha"` // 575489c9aed1698d11f1ed398e3c4cb3d45d2ca9
			Content   string `json:"Content"`   // Hello
			CreatedAt int64  `json:"CreatedAt"` // 1.642045112051e+12
			DepotID   int64  `json:"DepotId"`   // 2
			ID        int64  `json:"Id"`        // 330
			Index     int64  `json:"Index"`     // 1
			Path      string `json:"Path"`      // lib/src/extended_render_paragraph.dart
		} `json:"Comment"`
		RequestID string `json:"RequestId"` // db38b5e0-8f33-bf85-92fe-ca939d376006
	} `json:"Response"`
}

// CreateGitCommitComment 为某次提交创建一条评论
func (c *CommitClient) CreateGitCommitComment(req CreateGitCommitCommentReq) (resp CreateGitCommitCommentResp, err error) {
	err = c.Call(&req, &resp)
	return
}
