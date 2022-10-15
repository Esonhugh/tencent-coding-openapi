package baseClient

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"net/http"
)

const (
	MacOS_UserAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 Safari/605.1.15"
	CommonApiEndPoint = "https://e.coding.net/openapi"
	GetMeApiEndPoint  = "https://e.coding.net/api/me"
)

// ApiClientBase is an Object for Sending requests to the API
type ApiClientBase struct {
	// ApiEndPoint is the Coding Api EndPoints
	ApiEndPoint string

	// accessToken is variable for accessToken Store
	accessToken string

	// dataflow.Gout for Send Request Client
	*dataflow.Gout
}

// CreateApiClient is a function for creating a new ApiClientBase
func CreateApiClient(client ...*http.Client) *ApiClientBase {
	return &ApiClientBase{
		ApiEndPoint: CommonApiEndPoint,
		Gout:        gout.New(client...),
	}
}

// SetAccessToken is func to add Http Header Authorization: token
func (a *ApiClientBase) SetAccessToken(s string) *ApiClientBase {
	a.accessToken = s
	a.SetHeader(gout.H{
		"Authorization": fmt.Sprintf("token %s", s),
	})
	return a
}
