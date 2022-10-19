package cd

type DescribeCdApplicationReq struct {
	Action      string `json:"Action"`      // DescribeCdApplication
	Application string `json:"Application"` // test
}

func (req *DescribeCdApplicationReq) SetAction() string {
	req.Action = "DescribeCdApplication"
	return req.Action
}

type DescribeCdApplicationResp struct {
	Response struct {
		Data struct {
			ApplicationJsonContent string `json:"ApplicationJsonContent"` // {"cloudProviders":"kubernetes,hostserver","description":"","instancePort":80,"lastModifyNickName":"coding","name":"test","trafficGuards":[],"user":"coding"}
		} `json:"Data"`
		RequestID string `json:"RequestId"` // a95b0fcf-8e96-4a2c-5052-01184fa4caff
	} `json:"Response"`
}

// DescribeCdApplication 获取 CD 应用详情
func (c *CdClient) DescribeCdApplication(req DescribeCdApplicationReq) (resp DescribeCdApplicationResp, err error) {
	err = c.Call(&req, &resp)
	return
}
