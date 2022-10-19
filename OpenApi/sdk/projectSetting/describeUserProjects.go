package projectSetting

type DescribeUserProjectsReq struct {
	Action string `json:"Action"` // DescribeUserProjects
	UserID int64  `json:"userId"` // 2
}

func (req *DescribeUserProjectsReq) SetAction() string {
	req.Action = "DescribeUserProjects"
	return req.Action
}

type DescribeUserProjectsResp struct {
	Response struct {
		ProjectList []struct {
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
		} `json:"ProjectList"`
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// DescribeUserProjects 查询成员所在项目列表
func (p *ProjectSettingClient) DescribeUserProjects(req DescribeUserProjectsReq) (resp DescribeUserProjectsResp, err error) {
	err = p.Call(&req, &resp)
	return
}
