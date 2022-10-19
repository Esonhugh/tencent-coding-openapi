package ci

type DescribeCodingCIBuildLogRawReq struct {
	Action  string `json:"Action"`  // DescribeCodingCIBuildLogRaw
	BuildID int64  `json:"BuildId"` // 56
	Start   int64  `json:"Start"`   // 0
}

func (req *DescribeCodingCIBuildLogRawReq) SetAction() string {
	req.Action = "DescribeCodingCIBuildLogRaw"
	return req.Action
}

type DescribeCodingCIBuildLogRawResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeCodingCIBuildLogRaw 查询构建完整日志（原始日志 Raw）
func (c *CiClient) DescribeCodingCIBuildLogRaw(req DescribeCodingCIBuildLogRawReq) (resp DescribeCodingCIBuildLogRawResp, err error) {
	err = c.Call(&req, &resp)
	return
}
