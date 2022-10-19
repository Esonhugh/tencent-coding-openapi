package test

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/team"
	"testing"
)

func Test_Client(t *testing.T) {
	c := OpenApi.NewClient()
	meInfo, err := c.DescribeCodingProjects(team.DescribeCodingProjectsReq{
		PageSize:    10,
		PageNumber:  1,
		ProjectName: "",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(meInfo)
}
