package cd

type CancelCdPipelineReq struct {
	Action              string `json:"Action"`              // CancelCdPipeline
	PipelineExecutionID string `json:"PipelineExecutionId"` // 01FCWNV355ECJSZYXFQE3166VT
}

func (req *CancelCdPipelineReq) SetAction() string {
	req.Action = "CancelCdPipeline"
	return req.Action
}

type CancelCdPipelineResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // f78e5722-c4e6-09fc-664c-62e98a970fb0
	} `json:"Response"`
}

// CancelCdPipeline 取消执行中的 CD 部署流程
func (c *CdClient) CancelCdPipeline(req CancelCdPipelineReq) (resp CancelCdPipelineResp, err error) {
	err = c.Call(&req, &resp)
	return
}
