package TeamSDK

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type TeamSdkClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) TeamSdkClient {
	return TeamSdkClient{
		base,
	}
}

func (c *TeamSdkClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
