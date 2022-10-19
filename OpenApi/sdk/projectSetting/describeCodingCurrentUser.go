package projectSetting

type DescribeCodingCurrentUserReq struct {
	Action string `json:"Action"` // DescribeCodingCurrentUser
}

func (req *DescribeCodingCurrentUserReq) SetAction() string {
	req.Action = "DescribeCodingCurrentUser"
	return req.Action
}

type DescribeCodingCurrentUserResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 367bdaa8-e4c7-a070-ecf6-00fea8e6fab9
		User      struct {
			Avatar          string `json:"Avatar"`                 // https://coding-net-production-static-ci.codehub.cn/WM-TEXT-AVATAR-nfutKljCRlKcSLDTOmrv.jpg
			Email           string `json:"Email"`                  // test@coding.net
			EmailValidation int64  `json:"EmailValidation"`        // 1
			GlobalKey       string `json:"GlobalKey"`              // EHRIORBbfF
			ID              int64  `json:"Id"`                     // 150258
			Name            string `json:"Name"`                   // 张三
			NamePinYin      string `json:"NamePinYin"`             // zhangsan
			Phone           int64  `json:"Phone,string"`           // 13800138000
			PhoneRegionCode int64  `json:"PhoneRegionCode,string"` // +86
			PhoneValidation int64  `json:"PhoneValidation"`        // 1
			Status          int64  `json:"Status"`                 // 1
			TeamID          int64  `json:"TeamId"`                 // 102882
		} `json:"User"`
	} `json:"Response"`
}

// DescribeCodingCurrentUser 获取用户个人信息
func (p *ProjectSettingClient) DescribeCodingCurrentUser(req DescribeCodingCurrentUserReq) (resp DescribeCodingCurrentUserResp, err error) {
	err = p.Call(&req, &resp)
	return
}
