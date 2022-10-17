package CommonSDK

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"
)

type GetMeRequest struct {
}

type GetMeResponse struct {
	ID              int    `json:"id"`
	Avatar          string `json:"avatar"`
	CreateAt        int    `json:"created_at"`
	GlobalKey       string `json:"global_key"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	NamePinYin      string `json:"name_pinyin"`
	UpdateAt        int    `json:"update_at"`
	Path            string `json:"path"`
	Team            string `json:"team"`
	EmailValidation int    `json:"email_validation"`
}

func (c *CommonSdkClient) GetMe() (GetMeResponse, error) {
	var r GetMeResponse
	c.GET(baseClient.GetMeApiEndPoint).BindJSON(&r)
	return r, c.Do()
}
