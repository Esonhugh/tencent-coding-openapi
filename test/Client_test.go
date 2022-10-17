package test

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/ProjectSDK"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/TeamSDK"
	"testing"
)

func Test_Client(t *testing.T) {
	c := OpenApi.NewClient()
	meInfo, err := c.DescribeCodingProjects(TeamSDK.DescribeCodingProjectsRequest{
		PageSize:    10,
		PageNumber:  1,
		ProjectName: "",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(meInfo)
}

func Test_Client2(t *testing.T) {
	c := OpenApi.NewClient()
	meInfo, err := c.DescribeProjectByName(ProjectSDK.DescribeProjectByNameRequest{
		ProjectName: "srun4-portal-atour",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(meInfo)
}
