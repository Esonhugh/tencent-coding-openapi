package ci

type DescribeCodingCIBuildStepReq struct {
	Action  string `json:"Action"`  // DescribeCodingCIBuildStep
	BuildID int64  `json:"BuildId"` // 1
	StageID int64  `json:"StageId"` // 1
}

func (req *DescribeCodingCIBuildStepReq) SetAction() string {
	req.Action = "DescribeCodingCIBuildStep"
	return req.Action
}

type DescribeCodingCIBuildStepResp struct {
	Response struct {
		Data struct {
			StepJsonString string `json:"StepJsonString"` // [{"_class":"io.jenkins.blueocean.rest.impl.pipeline.PipelineStepImpl","_links":{"self":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-10-664474/runs/1/nodes/6/steps/7/"},"actions":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-10-664474/runs/1/nodes/6/steps/7/actions/"}},"actions":[{"_class":"io.jenkins.blueocean.service.embedded.rest.ActionProxiesImpl","_links":{"self":{"_class":"io.jenkins.blueocean.rest.hal.Link","href":"/blue/rest/organizations/jenkins/pipelines/cci-10-664474/runs/1/nodes/6/steps/7/log/"}},"_class":"org.jenkinsci.plugins.workflow.support.actions.LogStorageAction","urlName":"log"}],"displayDescription":null,"displayName":"Check out from version control","durationInMillis":4490,"id":"7","input":null,"result":"SUCCESS","startTime":"2020-05-08T17:45:21.407+0800","state":"FINISHED","type":"STEP"}]
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeCodingCIBuildStep 获取构建任务指定阶段的步骤
func (c *CiClient) DescribeCodingCIBuildStep(req DescribeCodingCIBuildStepReq) (resp DescribeCodingCIBuildStepResp, err error) {
	err = c.Call(&req, &resp)
	return
}
