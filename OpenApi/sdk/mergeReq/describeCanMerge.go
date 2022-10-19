package mergeReq

type DescribeCanMergeReq struct {
	Action  string `json:"Action"`  // DescribeCanMerge
	DepotID int64  `json:"DepotId"` // 809883
	Source  string `json:"Source"`  // dev/branch1
	Target  string `json:"Target"`  // release
}

func (req *DescribeCanMergeReq) SetAction() string {
	req.Action = "DescribeCanMerge"
	return req.Action
}

type DescribeCanMergeResp struct {
	Response struct {
		MergeStatus string `json:"MergeStatus"` // MERGEABLE
		RequestID   string `json:"RequestId"`   // f4135d6c-9bc6-2dd6-97f4-b2a850dd9713
	} `json:"Response"`
}

// DescribeCanMerge 获取分支的合并状态
func (m *MergeReqClient) DescribeCanMerge(req DescribeCanMergeReq) (resp DescribeCanMergeResp, err error) {
	err = m.Call(&req, &resp)
	return
}
