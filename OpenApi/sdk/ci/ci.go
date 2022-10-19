package ci

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type CiClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) CiClient {
	return CiClient{
		base,
	}
}

func (c *CiClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
