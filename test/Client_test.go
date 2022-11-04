package test

import (
	"fmt"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/team"
	"github.com/esonhugh/tencent-coding-openapi/utils/Print"
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
	var t2 Print.Table
	t2.Header = []string{"ID", "Name", "DisplayName", "Description"}
	for _, each := range meInfo.Response.Data.ProjectList {
		t2.Body = append(t2.Body, []string{fmt.Sprintf("%v", each.ID), each.Name, each.DisplayName, each.Description})
	}
	t2.Print("ProjectList")
}

func Test_client2(t *testing.T) {
	c := OpenApi.NewClient()
	resp, err := c.GetMe()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
