package branchProtection

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type BranchProtectionClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) BranchProtectionClient {
	return BranchProtectionClient{
		base,
	}
}

func (c *BranchProtectionClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
