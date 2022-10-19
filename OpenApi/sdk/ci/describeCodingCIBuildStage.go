package ci

type DescribeCodingCIBuildStageReq struct {
	Action  string `json:"Action"`  // DescribeCodingCIBuildStage
	BuildID int64  `json:"BuildId"` // 1
}

func (req *DescribeCodingCIBuildStageReq) SetAction() string {
	req.Action = "DescribeCodingCIBuildStage"
	return req.Action
}

type DescribeCodingCIBuildStageResp struct {
	Response struct {
		Data struct {
			StageJsonString string `json:"StageJsonString"` // [{"_class":"io.jenkins.blueocean.rest.impl.pipeline.PipelineNodeImpl","_links":{"self":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/6/"},"actions":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/6/actions/"},"steps":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/6/steps/"}},"actions":[],"displayDescription":null,"displayName":"检出","durationInMillis":5694,"id":"6","input":null,"result":"SUCCESS","startTime":"2020-05-10T12:21:02.862+0800","state":"FINISHED","type":"STAGE","causeOfBlockage":null,"edges":[{"_class":"io.jenkins.blueocean.rest.impl.pipeline.PipelineNodeImpl$EdgeImpl","id":"11","type":"STAGE"}],"firstParent":null,"restartable":true},{"_class":"io.jenkins.blueocean.rest.impl.pipeline.PipelineNodeImpl","_links":{"self":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/11/"},"actions":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/11/actions/"},"steps":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/11/steps/"}},"actions":[],"displayDescription":null,"displayName":"构建","durationInMillis":427,"id":"11","input":null,"result":"SUCCESS","startTime":"2020-05-10T12:21:08.611+0800","state":"FINISHED","type":"STAGE","causeOfBlockage":null,"edges":[{"_class":"io.jenkins.blueocean.rest.impl.pipeline.PipelineNodeImpl$EdgeImpl","id":"24","type":"STAGE"}],"firstParent":"6","restartable":true},{"_class":"io.jenkins.blueocean.rest.impl.pipeline.PipelineNodeImpl","_links":{"self":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/24/"},"actions":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/24/actions/"},"steps":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/24/steps/"}},"actions":[],"displayDescription":null,"displayName":"测试","durationInMillis":97,"id":"24","input":null,"result":"SUCCESS","startTime":"2020-05-10T12:21:09.095+0800","state":"FINISHED","type":"STAGE","causeOfBlockage":null,"edges":[{"_class":"io.jenkins.blueocean.rest.impl.pipeline.PipelineNodeImpl$EdgeImpl","id":"32","type":"STAGE"}],"firstParent":"11","restartable":true},{"_class":"io.jenkins.blueocean.rest.impl.pipeline.PipelineNodeImpl","_links":{"self":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/32/"},"actions":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/32/actions/"},"steps":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-2-605077/runs/1/nodes/32/steps/"}},"actions":[],"displayDescription":null,"displayName":"部署","durationInMillis":37,"id":"32","input":null,"result":"SUCCESS","startTime":"2020-05-10T12:21:09.253+0800","state":"FINISHED","type":"STAGE","causeOfBlockage":null,"edges":[],"firstParent":"24","restartable":true}]
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeCodingCIBuildStage 获取构建任务阶段
func (c *CiClient) DescribeCodingCIBuildStage(req DescribeCodingCIBuildStageReq) (resp DescribeCodingCIBuildStageResp, err error) {
	err = c.Call(&req, &resp)
	return
}
