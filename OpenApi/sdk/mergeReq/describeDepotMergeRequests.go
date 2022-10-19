package mergeReq

type DescribeDepotMergeRequestsReq struct {
	Action             string        `json:"Action"`             // DescribeDepotMergeRequests
	CreatorGlobalKeys  []string      `json:"CreatorGlobalKeys"`  // coding-coding
	DepotID            int64         `json:"DepotId"`            // 2
	IsSortDirectionAsc bool          `json:"IsSortDirectionAsc"` // true
	KeyWord            string        `json:"KeyWord"`            // modify
	Labels             []string      `json:"Labels"`             // label-1
	MergerGlobalKeys   []interface{} `json:"MergerGlobalKeys"`   // <nil>
	PageNumber         int64         `json:"PageNumber"`         // 1
	PageSize           int64         `json:"PageSize"`           // 10
	ReviewerGlobalKeys []interface{} `json:"ReviewerGlobalKeys"` // <nil>
	Sort               string        `json:"Sort"`               // action_at
	Status             string        `json:"Status"`             // open
	TargetBranches     []string      `json:"TargetBranches"`     // master
}

func (req *DescribeDepotMergeRequestsReq) SetAction() string {
	req.Action = "DescribeDepotMergeRequests"
	return req.Action
}

type DescribeDepotMergeRequestsResp struct {
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
				Labels                []string      `json:"Labels"`                // label-1
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
		RequestID string `json:"RequestId"` // f30370a3-6ce1-6207-34c8-8c4e462f7064
	} `json:"Response"`
}

// DescribeDepotMergeRequests 查询仓库下的合并请求列表
func (m *MergeReqClient) DescribeDepotMergeRequests(req DescribeDepotMergeRequestsReq) (resp DescribeDepotMergeRequestsResp, err error) {
	err = m.Call(&req, &resp)
	return
}
