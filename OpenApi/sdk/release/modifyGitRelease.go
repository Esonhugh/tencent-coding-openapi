package release

type ModifyGitReleaseReq struct {
	Action      string `json:"Action"`      // ModifyGitRelease
	DepotID     int64  `json:"DepotId"`     // 809883
	Description string `json:"Description"` // ModifyReleaseDescription
	Iid         int64  `json:"Iid"`         // 17
	Pre         bool   `json:"Pre"`         // true
	Title       string `json:"Title"`       // ModifyReleaseTitleName
}

func (req *ModifyGitReleaseReq) SetAction() string {
	req.Action = "ModifyGitRelease"
	return req.Action
}

type ModifyGitReleaseResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 542449d7-c975-9fbe-9389-7d009e742d0f
	} `json:"Response"`
}

// ModifyGitRelease 修改仓库的版本信息
func (r *ReleaseClient) ModifyGitRelease(req ModifyGitReleaseReq) (resp ModifyGitReleaseResp, err error) {
	err = r.Call(&req, &resp)
	return
}
