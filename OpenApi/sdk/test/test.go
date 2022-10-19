package test

import "github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"

type TestClient struct {
	*baseClient.ApiClientBase
}

func New(base *baseClient.ApiClientBase) TestClient {
	return TestClient{
		base,
	}
}

func (c *TestClient) ReInit(base *baseClient.ApiClientBase) {
	c.ApiClientBase = base
}
