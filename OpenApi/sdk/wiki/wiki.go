package wiki

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type WikiClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) WikiClient {
	return WikiClient{
		base,
	}
}

func (c *WikiClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
