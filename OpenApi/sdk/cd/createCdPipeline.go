package cd

type CreateCdPipelineReq struct {
	Action              string `json:"Action"`              // CreateCdPipeline
	PipelineJsonContent string `json:"PipelineJsonContent"` // {"keepWaitingPipelines":false,"limitConcurrent":true,"application":"test","codingNickname":"coding","roles":["团队管理员","团队所有者"],"lastModifiedBy":"Hello","name":"test","stages":[{"refId":"1","requisiteStageRefIds":[],"type":"wait","name":"等待","waitTime":10}],"index":0,"triggers":[],"updateTs":"1612317257852","relationProjects":[]}
}

func (req *CreateCdPipelineReq) SetAction() string {
	req.Action = "CreateCdPipeline"
	return req.Action
}

type CreateCdPipelineResp struct {
	Response struct {
		Data struct {
			TaskExecutionID  string `json:"TaskExecutionId"`  // 01EXRN2RDRPM9DW8HCHMMNKS1M
			TaskExecutionRef string `json:"TaskExecutionRef"` // /tasks/01EXRN2RDRPM9DW8HCHMMNKS1M
		} `json:"Data"`
		RequestID string `json:"RequestId"` // d6398a62-9af5-7ebd-13cb-07996a25917f
	} `json:"Response"`
}

// CreateCdPipeline 创建 CD 部署流程
func (c *CdClient) CreateCdPipeline(req CreateCdPipelineReq) (resp CreateCdPipelineResp, err error) {
	err = c.Call(&req, &resp)
	return
}
