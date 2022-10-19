package cd

type DescribeCdApplicationsReq struct {
	Action string `json:"Action"` // DescribeCdApplications
}

func (req *DescribeCdApplicationsReq) SetAction() string {
	req.Action = "DescribeCdApplications"
	return req.Action
}

type DescribeCdApplicationsResp struct {
	Response struct {
		Data struct {
			Applications []struct {
				Name string `json:"Name"` // dev
			} `json:"Applications"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // f7f613b5-83eb-438b-c196-9f7b4d01fde1
	} `json:"Response"`
}

// DescribeCdApplications 获取应用列表
func (c *CdClient) DescribeCdApplications(req DescribeCdApplicationsReq) (resp DescribeCdApplicationsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
