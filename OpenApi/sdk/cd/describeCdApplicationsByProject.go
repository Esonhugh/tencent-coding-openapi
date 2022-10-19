package cd

type DescribeCdApplicationsByProjectReq struct {
	Action      string `json:"Action"`      // DescribeCdApplicationsByProject
	ProjectName string `json:"ProjectName"` // test
}

func (req *DescribeCdApplicationsByProjectReq) SetAction() string {
	req.Action = "DescribeCdApplicationsByProject"
	return req.Action
}

type DescribeCdApplicationsByProjectResp struct {
	Response struct {
		Data struct {
			Applications []struct {
				Name string `json:"Name"` // dev
			} `json:"Applications"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 481be64b-f436-c025-2667-e89e7db69c6e
	} `json:"Response"`
}

// DescribeCdApplicationsByProject 根据项目名获取关联应用列表
func (c *CdClient) DescribeCdApplicationsByProject(req DescribeCdApplicationsByProjectReq) (resp DescribeCdApplicationsByProjectResp, err error) {
	err = c.Call(&req, &resp)
	return
}
