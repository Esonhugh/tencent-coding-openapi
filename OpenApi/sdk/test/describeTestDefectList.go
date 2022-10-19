package test

type DescribeTestDefectListReq struct {
	Action      string `json:"Action"`      // DescribeTestDefectList
	ProjectName string `json:"ProjectName"` // xx
	TestID      int64  `json:"TestId"`      // 1
}

func (req *DescribeTestDefectListReq) SetAction() string {
	req.Action = "DescribeTestDefectList"
	return req.Action
}

type DescribeTestDefectListResp struct {
	Response struct {
		Data struct {
			Defects []struct {
				AssignedTo  string `json:"AssignedTo"`  //
				Author      string `json:"Author"`      // coding
				CreatedAt   string `json:"CreatedAt"`   // 2021-07-09 15:13:46
				Description string `json:"Description"` //
				ID          int64  `json:"Id"`          // 1973
				Name        string `json:"Name"`        // 一个缺陷
				Status      string `json:"Status"`      // TODO
				StatusName  string `json:"StatusName"`  // 待处理
			} `json:"Defects"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 2efcd378-f6cf-83c0-b705-86896740b45e
	} `json:"Response"`
}

// DescribeTestDefectList 测试任务关联的缺陷列表
func (t *TestClient) DescribeTestDefectList(req DescribeTestDefectListReq) (resp DescribeTestDefectListResp, err error) {
	err = t.Call(&req, &resp)
	return
}
