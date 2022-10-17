package test

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"testing"
)

func Test_Client(t *testing.T) {
	c := OpenApi.NewClient()
	c.ApiEndPoint = "http://127.0.0.1:8080/api"
	c.SetToken(false, "114514")
	meInfo, err := c.GetMe()
	if err != nil {
		t.Error(err)
	}
	t.Log(meInfo)
}
