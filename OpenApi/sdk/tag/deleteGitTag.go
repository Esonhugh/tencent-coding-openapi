package tag

type DeleteGitTagReq struct {
	Action  string `json:"Action"`  // DeleteGitTag
	DepotID int64  `json:"DepotId"` // 1001
	TagName string `json:"TagName"` // tag-demo
}

func (req *DeleteGitTagReq) SetAction() string {
	req.Action = "DeleteGitTag"
	return req.Action
}

type DeleteGitTagResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // aa9dc0b9-6c8c-ed2c-ac54-191b40a63e81
	} `json:"Response"`
}

// DeleteGitTag 删除标签
func (t *TagClient) DeleteGitTag(req DeleteGitTagReq) (resp DeleteGitTagResp, err error) {
	err = t.Call(&req, &resp)
	return
}
