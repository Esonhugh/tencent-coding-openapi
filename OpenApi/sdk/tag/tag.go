package tag

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type TagClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) TagClient {
	return TagClient{
		base,
	}
}

func (c *TagClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
