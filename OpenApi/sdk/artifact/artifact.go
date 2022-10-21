package artifact

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"
)

type ArtifactClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) ArtifactClient {
	return ArtifactClient{
		base,
	}
}

func (c *ArtifactClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
