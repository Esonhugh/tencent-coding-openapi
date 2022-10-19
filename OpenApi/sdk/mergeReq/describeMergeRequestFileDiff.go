package mergeReq

type DescribeMergeRequestFileDiffReq struct {
	Action  string `json:"Action"`  // DescribeMergeRequestFileDiff
	DepotID int64  `json:"DepotId"` // 46
	MergeID int64  `json:"MergeId"` // 5
}

func (req *DescribeMergeRequestFileDiffReq) SetAction() string {
	req.Action = "DescribeMergeRequestFileDiff"
	return req.Action
}

type DescribeMergeRequestFileDiffResp struct {
	Response struct {
		MergeRequestFileDiff struct {
			Deletions int64 `json:"Deletions"` // 1
			FileDiffs []struct {
				ChangeType string `json:"ChangeType"` // ADD
				Deletions  int64  `json:"Deletions"`  // 0
				Insertions int64  `json:"Insertions"` // 1
				ObjectID   string `json:"ObjectId"`   // 4868155eaef0a678c55181286d89fd85c4261631
				Path       string `json:"Path"`       // demo.txt
			} `json:"FileDiffs"`
			Insertions int64  `json:"Insertions"` // 2
			NewSha     string `json:"NewSha"`     // 7c86071af40405abad16c90dcf6cb72abcf8184d
			OldSha     string `json:"OldSha"`     // b1c78397d7465a719b92ba41adac6e1c8bccfcdb
		} `json:"MergeRequestFileDiff"`
		RequestID string `json:"RequestId"` // 7dc24d8c-7dbc-3f62-0571-39c92a347ba4
	} `json:"Response"`
}

// DescribeMergeRequestFileDiff 获取合并请求文件修改记录
func (m *MergeReqClient) DescribeMergeRequestFileDiff(req DescribeMergeRequestFileDiffReq) (resp DescribeMergeRequestFileDiffResp, err error) {
	err = m.Call(&req, &resp)
	return
}
