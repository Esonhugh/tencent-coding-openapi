package release

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type ReleaseClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) ReleaseClient {
	return ReleaseClient{
		base,
	}
}

func (c *ReleaseClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
