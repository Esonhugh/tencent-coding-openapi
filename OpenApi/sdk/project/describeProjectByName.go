package project

type DescribeProjectByNameReq struct {
	Action      string `json:"Action"`      // DescribeProjectByName
	ProjectName string `json:"ProjectName"` // coding
}

func (req *DescribeProjectByNameReq) SetAction() string {
	req.Action = "DescribeProjectByName"
	return req.Action
}

type DescribeProjectByNameResp struct {
	Response struct {
		Project struct {
			Archived    bool   `json:"Archived"`    // false
			CreatedAt   int64  `json:"CreatedAt"`   // 1.572933083682e+12
			Description string `json:"Description"` // CODING 示例项目
			DisplayName string `json:"DisplayName"` // 示例项目
			EndDate     int64  `json:"EndDate"`     // 0
			Icon        string `json:"Icon"`        // https://dn-coding-net-production-pp.codehub.cn/79a8bcc4-d9cc-4061-940d-5b3bb31bf571.png
			ID          int64  `json:"Id"`          // 2
			IsDemo      bool   `json:"IsDemo"`      // true
			MaxMember   int64  `json:"MaxMember"`   // 0
			Name        string `json:"Name"`        // coding-demo
			StartDate   int64  `json:"StartDate"`   // 0
			Status      int64  `json:"Status"`      // 1
			TeamID      int64  `json:"TeamId"`      // 12
			TeamOwnerID int64  `json:"TeamOwnerId"` // 0
			Type        int64  `json:"Type"`        // 2
			UpdatedAt   int64  `json:"UpdatedAt"`   // 1.572933083682e+12
			UserOwnerID int64  `json:"UserOwnerId"` // 0
		} `json:"Project"`
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// DescribeProjectByName 查询项目根据项目名称
func (p *ProjectClient) DescribeProjectByName(req DescribeProjectByNameReq) (resp DescribeProjectByNameResp, err error) {
	err = p.Call(&req, &resp)
	return
}
