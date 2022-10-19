package mergeReq

type ModifyGitMergeRequestRebaseReq struct {
	Action  string `json:"Action"`  // ModifyGitMergeRequestRebase
	DepotID int64  `json:"DepotId"` // 2
	MergeID int64  `json:"MergeId"` // 2
}

func (req *ModifyGitMergeRequestRebaseReq) SetAction() string {
	req.Action = "ModifyGitMergeRequestRebase"
	return req.Action
}

type ModifyGitMergeRequestRebaseResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // fd409d06-9765-c2d5-28e7-bee9cfca19c2
	} `json:"Response"`
}

// ModifyGitMergeRequestRebase 合并请求中的源分支进行rebase操作
func (m *MergeReqClient) ModifyGitMergeRequestRebase(req ModifyGitMergeRequestRebaseReq) (resp ModifyGitMergeRequestRebaseResp, err error) {
	err = m.Call(&req, &resp)
	return
}
