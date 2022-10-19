package ci

type DeleteCodingCIBuildReq struct {
	Action  string `json:"Action"`  // DeleteCodingCIBuild
	BuildID int64  `json:"BuildId"` // 55
}

func (req *DeleteCodingCIBuildReq) SetAction() string {
	req.Action = "DeleteCodingCIBuild"
	return req.Action
}

type DeleteCodingCIBuildResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DeleteCodingCIBuild 删除构建
func (c *CiClient) DeleteCodingCIBuild(req DeleteCodingCIBuildReq) (resp DeleteCodingCIBuildResp, err error) {
	err = c.Call(&req, &resp)
	return
}
