package mergeReq

type DescribeSelfMergeRequestsReq struct {
	Action             string `json:"Action"`             // DescribeSelfMergeRequests
	IsSortDirectionAsc bool   `json:"IsSortDirectionAsc"` // true
	KeyWord            string `json:"KeyWord"`            // modify
	Label              string `json:"Label"`              // label-1
	MergerGlobalKey    string `json:"MergerGlobalKey"`    //
	ReviewerGlobalKey  string `json:"ReviewerGlobalKey"`  //
	Sort               string `json:"Sort"`               // action_at
	Status             string `json:"Status"`             // open
	TargetBranch       string `json:"TargetBranch"`       // master
}

func (req *DescribeSelfMergeRequestsReq) SetAction() string {
	req.Action = "DescribeSelfMergeRequests"
	return req.Action
}

type DescribeSelfMergeRequestsResp struct {
	Response struct {
		Data struct {
			List []struct {
				ActionAt     int64 `json:"ActionAt"` // 1.642142065e+12
				ActionAuthor struct {
					Avatar    string `json:"Avatar"`    //
					Email     string `json:"Email"`     //
					GlobalKey string `json:"GlobalKey"` //
					ID        int64  `json:"Id"`        // 0
					Name      string `json:"Name"`      //
					Status    string `json:"Status"`    // UNKNOWN
					TeamID    int64  `json:"TeamId"`    // 0
				} `json:"ActionAuthor"`
				Author struct {
					Avatar    string `json:"Avatar"`    // /static/fruit_avatar/Fruit-20.png
					Email     string `json:"Email"`     //
					GlobalKey string `json:"GlobalKey"` // coding-coding
					ID        int64  `json:"Id"`        // 2
					Name      string `json:"Name"`      // coding
					Status    string `json:"Status"`    // INACTIVE
					TeamID    int64  `json:"TeamId"`    // 0
				} `json:"Author"`
				CommentCount          int64         `json:"CommentCount"`          // 0
				CreatedAt             int64         `json:"CreatedAt"`             // 1.642140131e+12
				DepotID               int64         `json:"DepotId"`               // 2
				DesBranch             string        `json:"DesBranch"`             // master
				Granted               int64         `json:"Granted"`               // 0
				ID                    int64         `json:"Id"`                    // 1
				Labels                []interface{} `json:"Labels"`                // <nil>
				MergeID               int64         `json:"MergeId"`               // 1
				Path                  string        `json:"Path"`                  // /p/project-d/d/extended_text/git/merge/1
				Reminded              bool          `json:"Reminded"`              // false
				Reviewers             []interface{} `json:"Reviewers"`             // <nil>
				SrcBranch             string        `json:"SrcBranch"`             // dev
				Status                string        `json:"Status"`                // CANMERGE
				TargetBranchProtected bool          `json:"TargetBranchProtected"` // false
				Title                 string        `json:"Title"`                 // modify mr title
				UpdateAt              int64         `json:"UpdateAt"`              // 1.642152009e+12
			} `json:"List"`
			PageNumber int64 `json:"PageNumber"` // 1
			PageSize   int64 `json:"PageSize"`   // 10
			TotalPage  int64 `json:"TotalPage"`  // 1
			TotalRow   int64 `json:"TotalRow"`   // 1
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 0f569eca-7575-e6fe-d02a-8d67d8c8b6be
	} `json:"Response"`
}

// DescribeSelfMergeRequests 查询自己的合并请求列表
func (m *MergeReqClient) DescribeSelfMergeRequests(req DescribeSelfMergeRequestsReq) (resp DescribeSelfMergeRequestsResp, err error) {
	err = m.Call(&req, &resp)
	return
}
