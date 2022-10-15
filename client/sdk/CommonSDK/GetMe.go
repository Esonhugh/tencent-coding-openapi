package CommonSDK

import (
	"github.com/esonhugh/tencent-coding-openapi/client/baseClient"
)

type GetMeResponse struct {
	ID              string `json:"id"`
	Avatar          string `json:"avatar"`
	CreateAt        string `json:"created_at"`
	GlobalKey       string `json:"global_key"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	NamePinYin      string `json:"name_pinyin"`
	UpdateAt        string `json:"update_at"`
	Path            string `json:"path"`
	Team            string `json:"team"`
	EmailValidation string `json:"email_validation"`
}

func (c *CommonSdkClient) GetMe() (GetMeResponse, error) {
	var r GetMeResponse
	err := c.GET(baseClient.CommonApiEndPoint).
		BindJSON(&r).
		Do()
	return r, err
}
