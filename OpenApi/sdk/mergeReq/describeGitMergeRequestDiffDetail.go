package mergeReq

type DescribeGitMergeRequestDiffDetailReq struct {
	Action  string `json:"Action"`  // DescribeGitMergeRequestDiffDetail
	DepotID int64  `json:"DepotId"` // 2
	MergeID int64  `json:"MergeId"` // 1
}

func (req *DescribeGitMergeRequestDiffDetailReq) SetAction() string {
	req.Action = "DescribeGitMergeRequestDiffDetail"
	return req.Action
}

type DescribeGitMergeRequestDiffDetailResp struct {
	Response struct {
		Detail struct {
			ChangeType string `json:"ChangeType"` // add
			Content    string `json:"Content"`    // diff --git a/source/test.dart b/source/test.dart
			Deletions  int64  `json:"Deletions"`  // 0
			Insertions int64  `json:"Insertions"` // 1
			Lines      []struct {
				Index   int64  `json:"Index"`   // 0
				LeftNo  int64  `json:"LeftNo"`  // 0
				Prefix  string `json:"Prefix"`  //
				RightNo int64  `json:"RightNo"` // 0
				Text    string `json:"Text"`    // @@ -0,0 +1 @@

			} `json:"Lines"`
			NewMode int64  `json:"NewMode,string"` // 100644
			OldMode int64  `json:"OldMode,string"` // 0
			Path    string `json:"Path"`           // source/test.dart
		} `json:"Detail"`
		RequestID string `json:"RequestId"` // 10534565-3d16-c847-9a0d-2fc0752b80a3
	} `json:"Response"`
}

// DescribeGitMergeRequestDiffDetail 查询合并请求文件的 diff 详情
func (m *MergeReqClient) DescribeGitMergeRequestDiffDetail(req DescribeGitMergeRequestDiffDetailReq) (resp DescribeGitMergeRequestDiffDetailResp, err error) {
	err = m.Call(&req, &resp)
	return
}
