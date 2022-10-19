package release

type CreateGitReleaseReq struct {
	Action          string `json:"Action"`          // CreateGitRelease
	CommitSha       string `json:"CommitSha"`       // 48ddd600f7ff5e21a283105866de6a80fda708d7
	DepotID         int64  `json:"DepotId"`         // 809883
	Description     string `json:"Description"`     // Description
	Pre             bool   `json:"Pre"`             // true
	TagName         string `json:"TagName"`         // TagName
	TargetCommitish string `json:"TargetCommitish"` // master
	Title           string `json:"Title"`           // Title
	UserID          int64  `json:"UserId"`          // 1
}

func (req *CreateGitReleaseReq) SetAction() string {
	req.Action = "CreateGitRelease"
	return req.Action
}

type CreateGitReleaseResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ee4beb6c-92c6-df4e-c598-7cb73c2105be
	} `json:"Response"`
}

// CreateGitRelease 创建仓库版本
func (r *ReleaseClient) CreateGitRelease(req CreateGitReleaseReq) (resp CreateGitReleaseResp, err error) {
	err = r.Call(&req, &resp)
	return
}
