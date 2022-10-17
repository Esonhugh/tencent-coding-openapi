package CommonSDK

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"
)

type CommonSdkClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) CommonSdkClient {
	return CommonSdkClient{
		base,
	}
}

func (c *CommonSdkClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
