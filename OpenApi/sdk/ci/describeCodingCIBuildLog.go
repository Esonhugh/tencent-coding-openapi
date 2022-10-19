package ci

type DescribeCodingCIBuildLogReq struct {
	Action  string `json:"Action"`  // DescribeCodingCIBuildLog
	BuildID int64  `json:"BuildId"` // 1
	StageID int64  `json:"StageId"` // 1
	StepID  int64  `json:"StepId"`  // 1
	Start   int64  `json:"start"`   // 1
}

func (req *DescribeCodingCIBuildLogReq) SetAction() string {
	req.Action = "DescribeCodingCIBuildLog"
	return req.Action
}

type DescribeCodingCIBuildLogResp struct {
	Response struct {
		Data struct {
			Log           string `json:"Log"`           // Started by user coding
			MoreData      bool   `json:"MoreData"`      // false
			TextDelivered int64  `json:"TextDelivered"` // 3003
			TextSize      int64  `json:"TextSize"`      // 3003
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeCodingCIBuildLog 获取构建日志
func (c *CiClient) DescribeCodingCIBuildLog(req DescribeCodingCIBuildLogReq) (resp DescribeCodingCIBuildLogResp, err error) {
	err = c.Call(&req, &resp)
	return
}
