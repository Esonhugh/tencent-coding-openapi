package project

type DescribeProjectMembersReq struct {
	Action     string `json:"Action"`     // DescribeProjectMembers
	PageNumber int64  `json:"PageNumber"` // 1
	PageSize   int64  `json:"PageSize"`   // 10
	ProjectID  int64  `json:"ProjectId"`  // 1
	RoleID     int64  `json:"RoleId"`     // 1
}

func (req *DescribeProjectMembersReq) SetAction() string {
	req.Action = "DescribeProjectMembers"
	return req.Action
}

type DescribeProjectMembersResp struct {
	Response struct {
		Data struct {
			PageNumber     int64 `json:"PageNumber"` // 1
			PageSize       int64 `json:"PageSize"`   // 1
			ProjectMembers []struct {
				Avatar          string `json:"Avatar"`          // http://e.coding.net/static/fruit_avatar/Fruit-4.png
				Email           string `json:"Email"`           // blockuser@gmail.com
				EmailValidation int64  `json:"EmailValidation"` // 1
				GlobalKey       string `json:"GlobalKey"`       // GK
				ID              int64  `json:"Id"`              // 6
				Name            string `json:"Name"`            // blockuser
				NamePinYin      string `json:"NamePinYin"`      // blockuser
				Phone           int64  `json:"Phone,string"`    // 13800138006
				PhoneValidation int64  `json:"PhoneValidation"` // 1
				Roles           []struct {
					RoleID       int64  `json:"RoleId"`       // 1
					RoleType     string `json:"RoleType"`     // ProjectMember
					RoleTypeName string `json:"RoleTypeName"` // 开发
				} `json:"Roles"`
				Status int64 `json:"Status"` // -1
				TeamID int64 `json:"TeamId"` // 1
			} `json:"ProjectMembers"`
			TotalCount int64 `json:"TotalCount"` // 6
		} `json:"Data"`
		RequestID string `json:"RequestId"` // de0438d0-1e5e-098f-872a-c96863ed3510
	} `json:"Response"`
}

// DescribeProjectMembers 查询项目成员列表
func (p *ProjectClient) DescribeProjectMembers(req DescribeProjectMembersReq) (resp DescribeProjectMembersResp, err error) {
	err = p.Call(&req, &resp)
	return
}
