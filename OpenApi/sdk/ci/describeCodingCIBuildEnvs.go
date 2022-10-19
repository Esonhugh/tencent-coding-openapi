package ci

type DescribeCodingCIBuildEnvsReq struct {
	Action  string `json:"Action"`  // DescribeCodingCIBuildEnvs
	BuildID int64  `json:"BuildId"` // 1
	Type    string `json:"Type"`    // PARAM
}

func (req *DescribeCodingCIBuildEnvsReq) SetAction() string {
	req.Action = "DescribeCodingCIBuildEnvs"
	return req.Action
}

type DescribeCodingCIBuildEnvsResp struct {
	Response struct {
		EnvList []struct {
			Name      string `json:"Name"`      // TRIGGER_USER_NAME
			Sensitive bool   `json:"Sensitive"` // false
			Value     string `json:"Value"`     // coding
		} `json:"EnvList"`
		RequestID string `json:"RequestId"` // 09066b12-9027-1a52-d3bf-d1f30440f4c0
	} `json:"Response"`
}

// DescribeCodingCIBuildEnvs 获取构建计划环境变量
func (c *CiClient) DescribeCodingCIBuildEnvs(req DescribeCodingCIBuildEnvsReq) (resp DescribeCodingCIBuildEnvsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
