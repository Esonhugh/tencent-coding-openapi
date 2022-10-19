package work

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type WorkClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) WorkClient {
	return WorkClient{
		base,
	}
}

func (c *WorkClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
