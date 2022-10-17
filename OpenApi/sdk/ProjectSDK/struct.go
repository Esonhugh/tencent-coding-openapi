package ProjectSDK

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type ProjectSdkClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) ProjectSdkClient {
	return ProjectSdkClient{
		base,
	}
}

func (c *ProjectSdkClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
