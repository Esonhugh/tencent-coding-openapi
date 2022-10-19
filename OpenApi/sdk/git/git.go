package git

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type GitClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) GitClient {
	return GitClient{
		base,
	}
}

func (c *GitClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
