package tag

type CreateGitTagReq struct {
	Action     string `json:"Action"`     // CreateGitTag
	DepotID    int64  `json:"DepotId"`    // 1001
	StartPoint string `json:"StartPoint"` // master
	TagName    string `json:"TagName"`    // tag-demo
	Message    string `json:"message"`    // this is my create tag demo
}

func (req *CreateGitTagReq) SetAction() string {
	req.Action = "CreateGitTag"
	return req.Action
}

type CreateGitTagResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // b6445913-3b7c-9907-c6b3-a2d63abfd5c6
	} `json:"Response"`
}

// CreateGitTag 创建标签
func (t *TagClient) CreateGitTag(req CreateGitTagReq) (resp CreateGitTagResp, err error) {
	err = t.Call(&req, &resp)
	return
}
