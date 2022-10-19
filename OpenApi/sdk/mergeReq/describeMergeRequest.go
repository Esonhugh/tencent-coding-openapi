package mergeReq

type DescribeMergeRequestReq struct {
	Action  string `json:"Action"`  // DescribeMergeRequest
	DepotID int64  `json:"DepotId"` // 46
	MergeID int64  `json:"MergeId"` // 5
}

func (req *DescribeMergeRequestReq) SetAction() string {
	req.Action = "DescribeMergeRequest"
	return req.Action
}

type DescribeMergeRequestResp struct {
	Response struct {
		MergeRequestInfo struct {
			Describe     string `json:"Describe"`     // ## 改动说明
			SourceBranch string `json:"SourceBranch"` // dev
			Status       string `json:"Status"`       // ACCEPTED
			TargetBranch string `json:"TargetBranch"` // master
			Title        string `json:"Title"`        // 测试合并
		} `json:"MergeRequestInfo"`
		RequestID string `json:"RequestId"` // bcc8760c-5abf-ab67-f88c-f29befab2c8e
	} `json:"Response"`
}

// DescribeMergeRequest 查询合并请求详情
func (m *MergeReqClient) DescribeMergeRequest(req DescribeMergeRequestReq) (resp DescribeMergeRequestResp, err error) {
	err = m.Call(&req, &resp)
	return
}
