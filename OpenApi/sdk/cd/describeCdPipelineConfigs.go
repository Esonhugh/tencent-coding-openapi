package cd

type DescribeCdPipelineConfigsReq struct {
	Action      string `json:"Action"`      // DescribeCdPipelineConfigs
	Application string `json:"Application"` // test
}

func (req *DescribeCdPipelineConfigsReq) SetAction() string {
	req.Action = "DescribeCdPipelineConfigs"
	return req.Action
}

type DescribeCdPipelineConfigsResp struct {
	Response struct {
		Data struct {
			PipelineConfigs []struct {
				ID          string `json:"Id"`          // c2b76217-ba5e-4032-abbe-888765b73db4
				JsonContent string `json:"JsonContent"` // {"application":"test","codingNickname":"coding","id":"c2b76217-ba5e-4032-abbe-888765b73db4","index":0,"keepWaitingPipelines":false,"lastModifiedBy":"coding","limitConcurrent":true,"name":"hello","relationProjects":[],"roles":["团队管理员","团队所有者"],"stages":[{"name":"等待","refId":"1","requisiteStageRefIds":[],"type":"wait","waitTime":5}],"triggers":[],"updateTs":"1612514032427"}
				Name        string `json:"Name"`        // hello
			} `json:"PipelineConfigs"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // eb99f5ed-1a32-48f3-01aa-f44c2c96e92e
	} `json:"Response"`
}

// DescribeCdPipelineConfigs 获取 CD 应用下的所有部署流程配置
func (c *CdClient) DescribeCdPipelineConfigs(req DescribeCdPipelineConfigsReq) (resp DescribeCdPipelineConfigsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
