package cd

type DescribeCdPipelineReq struct {
	Action              string `json:"Action"`              // DescribeCdPipeline
	PipelineExecutionID string `json:"PipelineExecutionId"` // 01EXNXJVMK6EVRHB49CH7FG2BH
}

func (req *DescribeCdPipelineReq) SetAction() string {
	req.Action = "DescribeCdPipeline"
	return req.Action
}

type DescribeCdPipelineResp struct {
	Response struct {
		Data struct {
			Application                  string `json:"Application"`                  // test
			PipelineConfigID             string `json:"PipelineConfigId"`             // 99d2b0cf-3225-4aa3-9bea-6e66f4608e63
			PipelineExecutionJsonContent string `json:"PipelineExecutionJsonContent"` // {"application":"test","authentication":{"allowedAccounts":[],"user":"coding"},"buildTime":1612514226935,"canceled":false,"customName":"20210205-test-test","endTime":1612514237321,"id":"01EXRNCWPRS3Y0GSG4JX4N4CTE","initialConfig":{},"keepWaitingPipelines":false,"limitConcurrent":true,"name":"test","notifications":[],"origin":"api","pipelineConfigId":"99d2b0cf-3225-4aa3-9bea-6e66f4608e63","stages":[{"context":{"startTime":1612514227125,"waitTime":10},"endTime":1612514237268,"id":"01EXRNCWQQBBR0G5RKFQQWSVC5","name":"等待","outputs":{},"refId":"1","requisiteStageRefIds":[],"startTime":1612514227020,"status":"SUCCEEDED","tasks":[{"endTime":1612514237216,"id":"1","implementingClass":"com.netflix.spinnaker.orca.pipeline.tasks.WaitTask","loopEnd":false,"loopStart":false,"name":"wait","stageEnd":true,"stageStart":true,"startTime":1612514227072,"status":"SUCCEEDED"}],"type":"wait"}],"startTime":1612514226971,"status":"SUCCEEDED","systemNotifications":[],"trigger":{"artifacts":[],"codingNickname":"coding","dryRun":false,"enabled":false,"eventId":"340779f8-c7c5-4b93-a1f7-41c79793a4a0","executionId":"01EXRNCWPRS3Y0GSG4JX4N4CTE","notifications":[],"parameters":{},"rebake":false,"resolvedExpectedArtifacts":[],"strategy":false,"syncPipelineCacheEnable":false,"type":"manual","user":"coding"},"type":"PIPELINE"}
			PipelineExecutionStatus      string `json:"PipelineExecutionStatus"`      // SUCCEEDED
			PipelineName                 string `json:"PipelineName"`                 // test
		} `json:"Data"`
		RequestID string `json:"RequestId"` // eba6432e-abdf-e27b-ab5a-ddeeaa270716
	} `json:"Response"`
}

// DescribeCdPipeline 获取 CD 部署流程执行记录
func (c *CdClient) DescribeCdPipeline(req DescribeCdPipelineReq) (resp DescribeCdPipelineResp, err error) {
	err = c.Call(&req, &resp)
	return
}
