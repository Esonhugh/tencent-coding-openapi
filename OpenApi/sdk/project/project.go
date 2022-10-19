package project

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type ProjectClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) ProjectClient {
	return ProjectClient{
		base,
	}
}

func (c *ProjectClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
