package cd

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type CdClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) CdClient {
	return CdClient{
		base,
	}
}

func (c *CdClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
