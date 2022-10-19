package commit

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type CommitClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) CommitClient {
	return CommitClient{
		base,
	}
}

func (c *CommitClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
