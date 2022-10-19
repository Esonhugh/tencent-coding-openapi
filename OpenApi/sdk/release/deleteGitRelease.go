package release

type DeleteGitReleaseReq struct {
	Action  string `json:"Action"`  // DeleteGitRelease
	DepotID int64  `json:"DepotId"` // 809883
	TagName string `json:"TagName"` // TagName
}

func (req *DeleteGitReleaseReq) SetAction() string {
	req.Action = "DeleteGitRelease"
	return req.Action
}

type DeleteGitReleaseResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 542449d7-c975-9fbe-9389-7d009e742d0f
	} `json:"Response"`
}

// DeleteGitRelease 删除仓库的版本息
func (r *ReleaseClient) DeleteGitRelease(req DeleteGitReleaseReq) (resp DeleteGitReleaseResp, err error) {
	err = r.Call(&req, &resp)
	return
}
