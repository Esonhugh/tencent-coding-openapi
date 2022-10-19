package mergeReq

type CreateGitMergeReqReq struct {
	Action     string `json:"Action"`     // CreateGitMergeReq
	Content    string `json:"Content"`    // Hello
	DepotID    int64  `json:"DepotId"`    // 809883
	DestBranch string `json:"DestBranch"` // release
	SrcBranch  string `json:"SrcBranch"`  // dev/branch1
	Title      string `json:"Title"`      // Hello
}

func (req *CreateGitMergeReqReq) SetAction() string {
	req.Action = "CreateGitMergeReq"
	return req.Action
}

type CreateGitMergeReqResp struct {
	Response struct {
		MergeInfo struct {
			DepotID        int64 `json:"DepotId"`        // 809883
			MergeRequestID int64 `json:"MergeRequestId"` // 11
			ProjectID      int64 `json:"ProjectId"`      // 797199
		} `json:"MergeInfo"`
		RequestID string `json:"RequestId"` // 042b8230-87bc-d8aa-20d5-daa06fd319a6
	} `json:"Response"`
}

// CreateGitMergeReq 创建 merge 请求
func (m *MergeReqClient) CreateGitMergeReq(req CreateGitMergeReqReq) (resp CreateGitMergeReqResp, err error) {
	err = m.Call(&req, &resp)
	return
}
