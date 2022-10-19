package cd

type DeleteCdPipelineReq struct {
	Action       string `json:"Action"`       // DeleteCdPipeline
	Application  string `json:"Application"`  // test
	PipelineName string `json:"PipelineName"` // test
}

func (req *DeleteCdPipelineReq) SetAction() string {
	req.Action = "DeleteCdPipeline"
	return req.Action
}

type DeleteCdPipelineResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 01d43d0b-ed71-ed55-4704-900a0802ad1d
	} `json:"Response"`
}

// DeleteCdPipeline 删除 CD 部署流程
func (c *CdClient) DeleteCdPipeline(req DeleteCdPipelineReq) (resp DeleteCdPipelineResp, err error) {
	err = c.Call(&req, &resp)
	return
}
