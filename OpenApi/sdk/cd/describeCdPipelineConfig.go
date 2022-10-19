package cd

type DescribeCdPipelineConfigReq struct {
	Action       string `json:"Action"`       // DescribeCdPipelineConfig
	Application  string `json:"Application"`  // test
	PipelineName string `json:"PipelineName"` // test
}

func (req *DescribeCdPipelineConfigReq) SetAction() string {
	req.Action = "DescribeCdPipelineConfig"
	return req.Action
}

type DescribeCdPipelineConfigResp struct {
	Response struct {
		Data struct {
			PipelineConfig struct {
				ID          string `json:"Id"`          // 99d2b0cf-3225-4aa3-9bea-6e66f4608e63
				JsonContent string `json:"JsonContent"` // {"application":"test","codingNickname":"coding","id":"99d2b0cf-3225-4aa3-9bea-6e66f4608e63","index":1,"keepWaitingPipelines":false,"lastModifiedBy":"coding","limitConcurrent":true,"name":"test","relationProjects":[],"roles":["团队管理员","团队所有者"],"stages":[{"name":"等待","refId":"1","requisiteStageRefIds":[],"type":"wait","waitTime":10}],"triggers":[],"updateTs":"1612514145161"}
				Name        string `json:"Name"`        // test
			} `json:"PipelineConfig"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 01c09ae1-4154-f119-2089-482afe5b2166
	} `json:"Response"`
}

// DescribeCdPipelineConfig 根据名称获取 CD 部署流程配置
func (c *CdClient) DescribeCdPipelineConfig(req DescribeCdPipelineConfigReq) (resp DescribeCdPipelineConfigResp, err error) {
	err = c.Call(&req, &resp)
	return
}
