package cd

type DescribeCdTaskReq struct {
	Action          string `json:"Action"`          // DescribeCdTask
	TaskExecutionID string `json:"TaskExecutionId"` // 01EXNWQJFCBG9PE1GG1NPWPV4T
}

func (req *DescribeCdTaskReq) SetAction() string {
	req.Action = "DescribeCdTask"
	return req.Action
}

type DescribeCdTaskResp struct {
	Response struct {
		Data struct {
			Application              string `json:"Application"`              // test
			TaskExecutionJsonContent string `json:"TaskExecutionJsonContent"` // {"application":"test","buildTime":1612513747872,"endTime":1612513748575,"execution":{"application":"test","authentication":{"allowedAccounts":[],"user":"coding"},"buildTime":1612513747872,"canceled":false,"description":"Create Application: test","endTime":1612513748575,"id":"01EXRMY8X07T4Q2N65EGKCHZQ8","initialConfig":{},"keepWaitingPipelines":false,"limitConcurrent":false,"notifications":[],"origin":"unknown","stages":[{"context":{"application":{"cloudProviders":"kubernetes,hostserver","email":"coding@coding.com","instancePort":80,"lastModifyNickName":"coding","name":"testteam1","nickName":"coding","permissions":{"EXECUTE":["团队管理员team1","团队所有者team1"],"READ":["团队管理员team1","团队所有者team1"],"WRITE":["团队管理员team1","团队所有者team1"]}},"application.name":"testteam1","codingNickname":"coding","newState":{"cloudProviders":"kubernetes,hostserver","email":"coding@coding.com","instancePort":80,"lastModifyNickName":"coding","name":"testteam1","nickName":"coding","user":"coding"},"notification.type":"upsertapplication","previousState":{},"refId":"0","requisiteStageRefIds":[],"user":"coding"},"endTime":1612513748522,"id":"01EXRMY8X0XTMY25Y7Z26V2JBG","name":"createApplication","outputs":{},"refId":"0","requisiteStageRefIds":[],"startTime":1612513747953,"status":"SUCCEEDED","tasks":[{"endTime":1612513748470,"id":"1","implementingClass":"com.netflix.spinnaker.orca.applications.tasks.UpsertApplicationTask","loopEnd":false,"loopStart":false,"name":"createApplication","stageEnd":true,"stageStart":true,"startTime":1612513748005,"status":"SUCCEEDED"}],"type":"createApplication"}],"startTime":1612513747903,"status":"SUCCEEDED","systemNotifications":[],"trigger":{"artifacts":[],"codingNickname":"coding","dryRun":false,"notifications":[],"parameters":{},"rebake":false,"resolvedExpectedArtifacts":[],"strategy":false,"type":"manual","user":"coding"},"type":"ORCHESTRATION"},"id":"01EXRMY8X07T4Q2N65EGKCHZQ8","name":"Create Application: test","startTime":1612513747903,"status":"SUCCEEDED","steps":[{"endTime":1612513748470,"id":"1","implementingClass":"com.netflix.spinnaker.orca.applications.tasks.UpsertApplicationTask","loopEnd":false,"loopStart":false,"name":"createApplication","stageEnd":true,"stageStart":true,"startTime":1612513748005,"status":"SUCCEEDED"}],"variables":[{"key":"notification.type","value":"upsertapplication"},{"key":"requisiteStageRefIds","value":[]},{"key":"application","value":{"cloudProviders":"kubernetes,hostserver","email":"coding@coding.com","instancePort":80,"lastModifyNickName":"coding","name":"testteam1","nickName":"coding","permissions":{"EXECUTE":["团队管理员team1","团队所有者team1"],"READ":["团队管理员team1","团队所有者team1"],"WRITE":["团队管理员team1","团队所有者team1"]}}},{"key":"codingNickname","value":"coding"},{"key":"refId","value":"0"},{"key":"user","value":"coding"},{"key":"newState","value":{"cloudProviders":"kubernetes,hostserver","email":"coding@coding.com","instancePort":80,"lastModifyNickName":"coding","name":"testteam1","nickName":"coding","user":"coding"}},{"key":"application.name","value":"testteam1"},{"key":"previousState","value":{}}]}
			TaskExecutionStatus      string `json:"TaskExecutionStatus"`      // SUCCEEDED
		} `json:"Data"`
		RequestID string `json:"RequestId"` // c607f1b4-12b4-7b42-7ff3-09a1f5c4fff1
	} `json:"Response"`
}

// DescribeCdTask 获取 CD 任务执行记录
func (c *CdClient) DescribeCdTask(req DescribeCdTaskReq) (resp DescribeCdTaskResp, err error) {
	err = c.Call(&req, &resp)
	return
}
