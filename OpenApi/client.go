package OpenApi

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/CommonSDK"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/TeamSDK"
)

// Client struct is real Client contains pointers and Bind All Client Func.
type Client struct {
	*baseClient.ApiClientBase
	CommonSDK.CommonSdkClient
	TeamSDK.TeamSdkClient
}

// Useless SubSDKClient
type SubSDKClient interface {
	ReInit(*baseClient.ApiClientBase)
}

// NewClient func returns a real OpenApiClient
func NewClient() *Client {
	base := baseClient.CreateApiClient(nil)
	return &Client{
		base,
		CommonSDK.New(base),
		TeamSDK.New(base),
	}
}

// SetToken is map of c.SetAccessToken They are same.
func (c *Client) SetToken(isOAuth bool, token string) {
	c.SetAccessToken(isOAuth, token)
	// c.CommonSdkClient.SetAccessToken(isOAuth, token)
}
