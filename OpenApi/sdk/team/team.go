package team

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type TeamClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) TeamClient {
	return TeamClient{
		base,
	}
}

func (c *TeamClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
