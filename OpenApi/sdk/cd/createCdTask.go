package cd

type CreateCdTaskReq struct {
	Action          string `json:"Action"`          // CreateCdTask
	TaskJsonContent string `json:"TaskJsonContent"` // {"job":[{"type":"createApplication","application":{"cloudProviders":"kubernetes,hostserver","email":"coding@coding.com","instancePort":80,"nickName":"coding","name":"test"},"user":"coding"}],"application":"test","description":"Create Application: test"}
}

func (req *CreateCdTaskReq) SetAction() string {
	req.Action = "CreateCdTask"
	return req.Action
}

type CreateCdTaskResp struct {
	Response struct {
		Data struct {
			TaskExecutionID  string `json:"TaskExecutionId"`  // 01EXNWQJFCBG9PE1GG1NPWPV4T
			TaskExecutionRef string `json:"TaskExecutionRef"` // /tasks/01EXNWQJFCBG9PE1GG1NPWPV4T
		} `json:"Data"`
		RequestID string `json:"RequestId"` // cb0dfdd2-d6e4-36e8-385e-b91ecf8d1877
	} `json:"Response"`
}

// CreateCdTask 执行 CD 任务
func (c *CdClient) CreateCdTask(req CreateCdTaskReq) (resp CreateCdTaskResp, err error) {
	err = c.Call(&req, &resp)
	return
}
