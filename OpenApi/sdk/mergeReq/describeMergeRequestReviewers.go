package mergeReq

type DescribeMergeRequestReviewersReq struct {
	Action  string `json:"Action"`  // DescribeMergeRequestReviewers
	DepotID int64  `json:"DepotId"` // 1
	MergeID int64  `json:"MergeId"` // 1
}

func (req *DescribeMergeRequestReviewersReq) SetAction() string {
	req.Action = "DescribeMergeRequestReviewers"
	return req.Action
}

type DescribeMergeRequestReviewersResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 5cd3685f-be5d-4238-889d-7847a22a061e
		Reviewers []struct {
			Avatar    string `json:"Avatar"`    // /static/fruit_avatar/Fruit-20.png
			Email     string `json:"Email"`     //
			GlobalKey string `json:"GlobalKey"` // coding-user
			ID        int64  `json:"Id"`        // 4
			Name      string `json:"Name"`      // user
			Status    string `json:"Status"`    // INACTIVE
			TeamID    int64  `json:"TeamId"`    // 0
		} `json:"Reviewers"`
	} `json:"Response"`
}

// DescribeMergeRequestReviewers 查询合并请求评论者列表
func (m *MergeReqClient) DescribeMergeRequestReviewers(req DescribeMergeRequestReviewersReq) (resp DescribeMergeRequestReviewersResp, err error) {
	err = m.Call(&req, &resp)
	return
}
