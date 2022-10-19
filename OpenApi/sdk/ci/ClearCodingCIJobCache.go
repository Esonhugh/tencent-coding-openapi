package ci

type ClearCodingCIJobCacheReq struct {
	ID int64 `json:"Id"` // 0
}

func (req *ClearCodingCIJobCacheReq) SetAction() string {
	return "ClearCodingCIJobCache"
}

type ClearCodingCIJobCacheResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ClearCodingCIJobCache 清理构建计划缓存
func (c *CiClient) ClearCodingCIJobCache(req ClearCodingCIJobCacheReq) (resp ClearCodingCIJobCacheResp, err error) {
	err = c.Call(&req, &resp)
	return
}
