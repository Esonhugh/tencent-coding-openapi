package team

type DescribeCodingProjectsReq struct {
	Action      string `json:"Action"`      // DescribeCodingProjects
	PageNumber  int64  `json:"PageNumber"`  // 1
	PageSize    int64  `json:"PageSize"`    // 10
	ProjectName string `json:"ProjectName"` // coding
}

func (req *DescribeCodingProjectsReq) SetAction() string {
	req.Action = "DescribeCodingProjects"
	return req.Action
}

type DescribeCodingProjectsResp struct {
	Response struct {
		Data struct {
			PageNumber  int64 `json:"PageNumber"` // 1
			PageSize    int64 `json:"PageSize"`   // 1
			ProjectList []struct {
				Archived    bool   `json:"Archived"`    // false
				CreatedAt   int64  `json:"CreatedAt"`   // 1.619580482e+12
				Description string `json:"Description"` //
				DisplayName string `json:"DisplayName"` // empty
				EndDate     int64  `json:"EndDate"`     // 0
				Icon        string `json:"Icon"`        // https://e.coding.net/static/project_icon/scenery-version-2-4.svg
				ID          int64  `json:"Id"`          // 1
				IsDemo      bool   `json:"IsDemo"`      // false
				MaxMember   int64  `json:"MaxMember"`   // 0
				Name        string `json:"Name"`        // empty
				StartDate   int64  `json:"StartDate"`   // 0
				Status      int64  `json:"Status"`      // 1
				TeamID      int64  `json:"TeamId"`      // 1
				TeamOwnerID int64  `json:"TeamOwnerId"` // 1
				Type        int64  `json:"Type"`        // 2
				UpdatedAt   int64  `json:"UpdatedAt"`   // 1.619580482e+12
				UserOwnerID int64  `json:"UserOwnerId"` // 0
			} `json:"ProjectList"`
			TotalCount int64 `json:"TotalCount"` // 1
		} `json:"Data"`
		RequestID string `json:"RequestId"` // f41034cf-169d-4373-0514-334885a6a9db
	} `json:"Response"`
}

// DescribeCodingProjects 查询团队内所有项目列表
func (t *TeamClient) DescribeCodingProjects(req DescribeCodingProjectsReq) (resp DescribeCodingProjectsResp, err error) {
	err = t.Call(&req, &resp)
	return
}
