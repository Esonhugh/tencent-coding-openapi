package team

type DescribeTeamMembersReq struct {
	Action     string `json:"Action"`     // DescribeTeamMembers
	PageNumber int64  `json:"PageNumber"` // 1
	PageSize   int64  `json:"PageSize"`   // 10
}

func (req *DescribeTeamMembersReq) SetAction() string {
	req.Action = "DescribeTeamMembers"
	return req.Action
}

type DescribeTeamMembersResp struct {
	Response struct {
		Data struct {
			PageNumber  int64 `json:"PageNumber"` // 1
			PageSize    int64 `json:"PageSize"`   // 1
			TeamMembers []struct {
				Avatar          string `json:"Avatar"`          // http://e.coding.net/static/fruit_avatar/Fruit-4.png
				Email           string `json:"Email"`           // blockuser@gmail.com
				EmailValidation int64  `json:"EmailValidation"` // 1
				ID              int64  `json:"Id"`              // 6
				Name            string `json:"Name"`            // blockuser
				NamePinYin      string `json:"NamePinYin"`      // blockuser
				Phone           int64  `json:"Phone,string"`    // 13800138006
				PhoneValidation int64  `json:"PhoneValidation"` // 1
				Status          int64  `json:"Status"`          // -1
				TeamID          int64  `json:"TeamId"`          // 1
			} `json:"TeamMembers"`
			TotalCount int64 `json:"TotalCount"` // 6
		} `json:"Data"`
		RequestID string `json:"RequestId"` // de0438d0-1e5e-098f-872a-c96863ed3510
	} `json:"Response"`
}

// DescribeTeamMembers 查询团队成员列表
func (t *TeamClient) DescribeTeamMembers(req DescribeTeamMembersReq) (resp DescribeTeamMembersResp, err error) {
	err = t.Call(&req, &resp)
	return
}
