package branch

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type BranchClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) BranchClient {
	return BranchClient{
		base,
	}
}

func (c *BranchClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
