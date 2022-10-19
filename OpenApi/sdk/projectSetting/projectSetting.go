package projectSetting

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type ProjectSettingClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) ProjectSettingClient {
	return ProjectSettingClient{
		base,
	}
}

func (c *ProjectSettingClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
