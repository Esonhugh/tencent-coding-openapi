package ProjectSDK

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"
)

type DescribeOneProjectRequest struct {
	Action    string `json:"Action"`
	ProjectId int    `json:"ProjectId"`
}

func (req *DescribeOneProjectRequest) SetAction() string {
	req.Action = "DescribeOneProjectRequest"
	return req.Action
}

type DescribeOneProjectResponse struct {
	Response struct {
		RequestId string                   `json:"RequestId"`
		Project   baseClient.ProjectObject `json:"Project"`
	} `json:"Response"`
}

func (c *ProjectSdkClient) DescribeOneProject(req DescribeOneProjectRequest) (rsp DescribeOneProjectResponse, e error) {
	e = c.CallAction(&req, &rsp)
	return
}
