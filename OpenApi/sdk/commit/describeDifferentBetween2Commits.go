package commit

type DescribeDifferentBetween2CommitsReq struct {
	Action  string `json:"Action"`  // DescribeDifferentBetween2Commits
	DepotID int64  `json:"DepotId"` // 5001
	Path    string `json:"Path"`    // root/test.java
	Source  string `json:"Source"`  // master
	Target  string `json:"Target"`  // dev
}

func (req *DescribeDifferentBetween2CommitsReq) SetAction() string {
	req.Action = "DescribeDifferentBetween2Commits"
	return req.Action
}

type DescribeDifferentBetween2CommitsResp struct {
	Response struct {
		DiffFileInfo struct {
			Deletions      int64 `json:"Deletions"` // 0
			DifferentLines []struct {
				Index   int64  `json:"Index"`   // 0
				LeftNo  int64  `json:"LeftNo"`  // 0
				Prefix  string `json:"Prefix"`  //
				RightNo int64  `json:"RightNo"` // 0
				Text    string `json:"Text"`    // @@ -0,0 +1 @@

			} `json:"DifferentLines"`
			Insertions int64 `json:"Insertions"` // 1
		} `json:"DiffFileInfo"`
		RequestID string `json:"RequestId"` // 8122c513-7ad1-878c-6187-a764b4616a79
	} `json:"Response"`
}

// DescribeDifferentBetween2Commits 两次提交之间的文件差异
func (c *CommitClient) DescribeDifferentBetween2Commits(req DescribeDifferentBetween2CommitsReq) (resp DescribeDifferentBetween2CommitsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
