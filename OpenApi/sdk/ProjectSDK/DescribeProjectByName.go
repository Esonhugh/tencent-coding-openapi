package ProjectSDK

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"
)

type DescribeProjectByNameRequest struct {
	Action      string `json:"Action"`
	ProjectName string `json:"ProjectName"`
}

func (req *DescribeProjectByNameRequest) SetAction() string {
	req.Action = "DescribeProjectByName"
	return req.Action
}

type DescribeProjectByNameResponse struct {
	Response struct {
		RequestId string                   `json:"RequestId"`
		Project   baseClient.ProjectObject `json:"Project"`
	}
}

func (c *ProjectSdkClient) DescribeProjectByName(request DescribeProjectByNameRequest) (response DescribeProjectByNameResponse, err error) {
	// err = c.CallAction("DescribeProjectByName", &request, &response)
	err = c.Call(&request, &response)
	return
}
