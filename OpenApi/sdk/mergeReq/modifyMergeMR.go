package mergeReq

type ModifyMergeMRReq struct {
	Action            string `json:"Action"`            // ModifyMergeMR
	DepotID           int64  `json:"DepotId"`           // 809883
	IsDelSourceBranch bool   `json:"IsDelSourceBranch"` // true
	IsFastForward     bool   `json:"IsFastForward"`     // false
	MergeID           int64  `json:"MergeId"`           // 14
	Message           string `json:"Message"`           // Hello
	Squash            bool   `json:"Squash"`            // false
}

func (req *ModifyMergeMRReq) SetAction() string {
	req.Action = "ModifyMergeMR"
	return req.Action
}

type ModifyMergeMRResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 5f988d01-e86c-53c6-1170-c2b67eb9d335
	} `json:"Response"`
}

// ModifyMergeMR 合入合并请求
func (m *MergeReqClient) ModifyMergeMR(req ModifyMergeMRReq) (resp ModifyMergeMRResp, err error) {
	err = m.Call(&req, &resp)
	return
}
