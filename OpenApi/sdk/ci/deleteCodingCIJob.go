package ci

type DeleteCodingCIJobReq struct {
	Action  string `json:"Action"`  // DescribeCodingCIJob
	Version string `json:"Version"` // 1
	Region  string `json:"Region"`
	JobId   int    `json:"JobId"`
}

type DeleteCodingCIJobResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

func (req *DeleteCodingCIJobReq) SetAction() string {
	req.Action = "DeleteCodingCIJob"
	return req.Action
}

// DeleteCodingCIJob 删除构建计划
func (c *CiClient) DeleteCodingCIJob(req DeleteCodingCIJobReq) (resp DeleteCodingCIJobResp, err error) {
	err = c.Call(&req, &resp)
	return
}
