package mergeReq

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type MergeReqClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) MergeReqClient {
	return MergeReqClient{
		base,
	}
}

func (c *MergeReqClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
