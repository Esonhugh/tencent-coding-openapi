package ci

type StopCodingCIBuildReq struct {
	Action  string `json:"Action"`  // StopCodingCIBuild
	BuildID int64  `json:"BuildId"` // 1
}

func (req *StopCodingCIBuildReq) SetAction() string {
	req.Action = "StopCodingCIBuild"
	return req.Action
}

type StopCodingCIBuildResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// StopCodingCIBuild 停止构建
func (c *CiClient) StopCodingCIBuild(req StopCodingCIBuildReq) (resp StopCodingCIBuildResp, err error) {
	err = c.Call(&req, &resp)
	return
}
