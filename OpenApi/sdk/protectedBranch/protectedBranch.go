package protectedBranch

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type ProtectedBranchClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) ProtectedBranchClient {
	return ProtectedBranchClient{
		base,
	}
}

func (c *ProtectedBranchClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
