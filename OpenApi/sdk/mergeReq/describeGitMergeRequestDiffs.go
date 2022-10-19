package mergeReq

type DescribeGitMergeRequestDiffsReq struct {
	Action  string `json:"Action"`  // DescribeGitMergeRequestDiffs
	DepotID int64  `json:"DepotId"` // 2
	MergeID int64  `json:"MergeId"` // 2
}

func (req *DescribeGitMergeRequestDiffsReq) SetAction() string {
	req.Action = "DescribeGitMergeRequestDiffs"
	return req.Action
}

type DescribeGitMergeRequestDiffsResp struct {
	Response struct {
		Diff struct {
			Deletions  int64  `json:"Deletions"`  // 0
			Insertions int64  `json:"Insertions"` // 1
			IsLarge    bool   `json:"IsLarge"`    // false
			NewSha     string `json:"NewSha"`     // e06ebced0f890db2265b8362b3e3f24f30374413
			OldSha     string `json:"OldSha"`     // 125e0345107601d7e3d451032e6a40ecf0b4c548
			Paths      []struct {
				BlobSha    string `json:"BlobSha"`    // 81f4282e8e0bcd8f765ca96b7ef2b8b9050f4a97
				ChangeType string `json:"ChangeType"` // ADD
				Deletions  int64  `json:"Deletions"`  // 0
				Insertions int64  `json:"Insertions"` // 1
				Path       string `json:"Path"`       // src/qw.js
				Size       int64  `json:"Size"`       // 2
			} `json:"Paths"`
		} `json:"Diff"`
		RequestID string `json:"RequestId"` // 7171c2a8-4bb1-d88b-470d-e921bdb18188
	} `json:"Response"`
}

// DescribeGitMergeRequestDiffs 查询合并请求 diff 信息的文件列表
func (m *MergeReqClient) DescribeGitMergeRequestDiffs(req DescribeGitMergeRequestDiffsReq) (resp DescribeGitMergeRequestDiffsResp, err error) {
	err = m.Call(&req, &resp)
	return
}
