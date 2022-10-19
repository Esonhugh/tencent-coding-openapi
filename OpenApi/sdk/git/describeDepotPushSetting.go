package git

type DescribeDepotPushSettingReq struct {
	Action    string `json:"Action"`    // DescribeDepotPushSetting
	DepotPath string `json:"DepotPath"` // codingcorp/project-d/depot-2
}

func (req *DescribeDepotPushSettingReq) SetAction() string {
	req.Action = "DescribeDepotPushSetting"
	return req.Action
}

type DescribeDepotPushSettingResp struct {
	Response struct {
		Data struct {
			CheckCommitAuthor           bool   `json:"CheckCommitAuthor"`           // false
			CommitMessageMustMatchRegex string `json:"CommitMessageMustMatchRegex"` //
			DenyForcePush               bool   `json:"DenyForcePush"`               // false
			PushDenyFile                string `json:"PushDenyFile"`                //
			PushFileSize                string `json:"PushFileSize"`                //
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 70095d4a-8716-486b-b279-a778722060a8
	} `json:"Response"`
}

// DescribeDepotPushSetting 查询仓库推送设置
func (g *GitClient) DescribeDepotPushSetting(req DescribeDepotPushSettingReq) (resp DescribeDepotPushSettingResp, err error) {
	err = g.Call(&req, &resp)
	return
}
