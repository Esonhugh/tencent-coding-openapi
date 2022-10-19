package mergeReq

type ModifyGitMergeRequestReq struct {
	Action  string `json:"Action"`  // ModifyGitMergeRequest
	Content string `json:"Content"` // modify mr content
	DepotID int64  `json:"DepotId"` // 2
	MergeID int64  `json:"MergeId"` // 1
	Title   string `json:"Title"`   // modify mr title
}

func (req *ModifyGitMergeRequestReq) SetAction() string {
	req.Action = "ModifyGitMergeRequest"
	return req.Action
}

type ModifyGitMergeRequestResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // b275dc2e-792e-d81b-2b61-7dc8e60aeb41
	} `json:"Response"`
}

// ModifyGitMergeRequest 修改合并请求的标题和描述信息
func (m *MergeReqClient) ModifyGitMergeRequest(req ModifyGitMergeRequestReq) (resp ModifyGitMergeRequestResp, err error) {
	err = m.Call(&req, &resp)
	return
}
