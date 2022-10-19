package git

type DescribeDepotSpecsReq struct {
	Action string `json:"Action"` // DescribeDepotSpecs
}

func (req *DescribeDepotSpecsReq) SetAction() string {
	req.Action = "DescribeDepotSpecs"
	return req.Action
}

type DescribeDepotSpecsResp struct {
	Response struct {
		DepotSpecs []struct {
			Description string `json:"Description"` // 支持 master 分支
			Name        string `json:"Name"`        // 单分支规范
			Type        string `json:"Type"`        // system
		} `json:"DepotSpecs"`
		RequestID string `json:"RequestId"` // 719f32c7-4b52-4836-9b0c-12122f19bde3
	} `json:"Response"`
}

// DescribeDepotSpecs 查询仓库规范列表
func (g *GitClient) DescribeDepotSpecs(req DescribeDepotSpecsReq) (resp DescribeDepotSpecsResp, err error) {
	err = g.Call(&req, &resp)
	return
}
