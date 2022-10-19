package cd

type ModifyCdPipelineReq struct {
	Action              string `json:"Action"`              // ModifyCdPipeline
	PipelineConfigID    string `json:"PipelineConfigId"`    // c2b76217-ba5e-4032-abbe-888765b73db4
	PipelineJsonContent string `json:"PipelineJsonContent"` // {"keepWaitingPipelines":false,"lastModifiedBy":"coding","roles":["团队管理员","团队所有者"],"index":0,"triggers":[],"relationProjects":[],"limitConcurrent":true,"application":"test","codingNickname":"coding","name":"hello","stages":[{"name":"等待","refId":"1","requisiteStageRefIds":[],"type":"wait","waitTime":5}],"id":"c2b76217-ba5e-4032-abbe-888765b73db4","updateTs":"1612421444090"}
}

func (req *ModifyCdPipelineReq) SetAction() string {
	req.Action = "ModifyCdPipeline"
	return req.Action
}

type ModifyCdPipelineResp struct {
	Response struct {
		Data struct {
			TaskExecutionID  string `json:"TaskExecutionId"`  // 01EXRN6M97M91Q3E00QFMC4FZ8
			TaskExecutionRef string `json:"TaskExecutionRef"` // /tasks/01EXRN6M97M91Q3E00QFMC4FZ8
		} `json:"Data"`
		RequestID string `json:"RequestId"` // c0d671f1-c998-dda9-b2aa-8276cb72b619
	} `json:"Response"`
}

// ModifyCdPipeline 修改 CD 部署流程
func (c *CdClient) ModifyCdPipeline(req ModifyCdPipelineReq) (resp ModifyCdPipelineResp, err error) {
	err = c.Call(&req, &resp)
	return
}
