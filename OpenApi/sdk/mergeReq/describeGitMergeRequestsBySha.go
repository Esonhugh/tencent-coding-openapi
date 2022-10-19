package mergeReq

type DescribeGitMergeRequestsByShaReq struct {
	Action  string `json:"Action"`  // DescribeGitMergeRequestsBySha
	DepotID int64  `json:"DepotId"` // 2
	Sha     string `json:"Sha"`     // a9891701d0cfea86339cae013c1401ef0ed8edef
}

func (req *DescribeGitMergeRequestsByShaReq) SetAction() string {
	req.Action = "DescribeGitMergeRequestsBySha"
	return req.Action
}

type DescribeGitMergeRequestsByShaResp struct {
	Response struct {
		Details []struct {
			ActionAt     int64 `json:"ActionAt"` // 1.642140131e+12
			ActionAuthor struct {
				Avatar    string `json:"Avatar"`    //
				Email     string `json:"Email"`     //
				GlobalKey string `json:"GlobalKey"` //
				ID        int64  `json:"Id"`        // 0
				Name      string `json:"Name"`      //
				Status    string `json:"Status"`    //
				TeamID    int64  `json:"TeamId"`    // 0
			} `json:"ActionAuthor"`
			Author struct {
				Avatar    string `json:"Avatar"`    // /static/fruit_avatar/Fruit-20.png
				Email     string `json:"Email"`     // coding@coding.com
				GlobalKey string `json:"GlobalKey"` // coding-coding
				ID        int64  `json:"Id"`        // 2
				Name      string `json:"Name"`      // coding
				Status    string `json:"Status"`    // INACTIVE
				TeamID    int64  `json:"TeamId"`    // 1
			} `json:"Author"`
			Content      string `json:"Content"`      // test merge description
			ContentHtml  string `json:"ContentHtml"`  // <p>test merge description</p>
			CreateAt     int64  `json:"CreateAt"`     // 1.642140131e+12
			DepotID      int64  `json:"DepotId"`      // 2
			ID           int64  `json:"Id"`           // 1
			MergeID      int64  `json:"MergeId"`      // 1
			MergeStatus  string `json:"MergeStatus"`  // CANMERGE
			MergedSha    string `json:"MergedSha"`    //
			SourceBranch string `json:"SourceBranch"` // dev
			SourceSha    string `json:"SourceSha"`    // a9891701d0cfea86339cae013c1401ef0ed8edef
			TargetBranch string `json:"TargetBranch"` // master
			TargetSha    string `json:"TargetSha"`    // 6d19499a8f91da496c79acfa142f9a2b747ea76e
			Title        string `json:"Title"`        // test merge title
			UpdateAt     int64  `json:"UpdateAt"`     // 1.642140131e+12
		} `json:"Details"`
		RequestID string `json:"RequestId"` // ceca62de-5742-08f1-91f5-05f103fd228b
	} `json:"Response"`
}

// DescribeGitMergeRequestsBySha 查询含有某次提交的合并请求
func (m *MergeReqClient) DescribeGitMergeRequestsBySha(req DescribeGitMergeRequestsByShaReq) (resp DescribeGitMergeRequestsByShaResp, err error) {
	err = m.Call(&req, &resp)
	return
}
