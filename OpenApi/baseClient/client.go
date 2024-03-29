package baseClient

import (
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"net/http"
)

const (
	MacOS_UserAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 Safari/605.1.15"
	CommonApiEndPoint = "https://e.coding.net/open-api"
	// CommonApiEndPoint = "http://127.0.0.1:8080"
	GetMeApiEndPoint = "https://e.coding.net/api/me"
)

// ApiClientBase is an Object for Sending requests to the API. The basic implementation of APIClient internal.
type ApiClientBase struct {
	// ApiEndPoint is the Coding Api EndPoints
	ApiEndPoint string

	// accessToken is variable of user's accessToken
	accessToken string

	// bearerToken is variable for OAuthed Client AccessToken.
	bearerToken string

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
func (a *ApiClientBase) SetAccessToken(isOAuth bool, s string) *ApiClientBase {
	if isOAuth {
		a.bearerToken = s
	} else {
		a.accessToken = s
	}
	return a
}

// setAuthHeader is func internal and set Header when do.
func (a *ApiClientBase) setAuthHeader() *ApiClientBase {
	if a.bearerToken != "" {
		a.SetHeader(gout.H{
			"Authorization": "Bearer " + a.accessToken,
		})
	}
	if a.accessToken != "" {
		a.SetHeader(gout.H{
			"Authorization": "token " + a.accessToken,
		})
	}
	return a
}
