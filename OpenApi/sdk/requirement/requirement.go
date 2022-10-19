package requirement

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type RequirementClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) RequirementClient {
	return RequirementClient{
		base,
	}
}

func (c *RequirementClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
