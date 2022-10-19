package issue

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type IssueClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) IssueClient {
	return IssueClient{
		base,
	}
}

func (c *IssueClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
