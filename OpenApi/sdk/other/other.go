package other

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"
)

type OtherClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) OtherClient {
	return OtherClient{
		base,
	}
}

func (c *OtherClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
