package iteration

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type IterationClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) IterationClient {
	return IterationClient{
		base,
	}
}

func (c *IterationClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
