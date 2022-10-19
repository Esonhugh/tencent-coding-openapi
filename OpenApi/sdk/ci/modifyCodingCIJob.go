package ci

type ModifyCodingCIJobReq struct {
	AutoCancelSameMergeRequest bool   `json:"AutoCancelSameMergeRequest"` // true
	AutoCancelSameRevision     bool   `json:"AutoCancelSameRevision"`     // true
	BranchRegex                string `json:"BranchRegex"`                //
	BranchSelector             string `json:"BranchSelector"`             //
	DepotID                    int64  `json:"DepotId"`                    // 0
	DepotType                  string `json:"DepotType"`                  // NONE
	ExecuteIn                  string `json:"ExecuteIn"`                  // CVM
	HookType                   string `json:"HookType"`                   // DEFAULT
	ID                         int64  `json:"Id"`                         // 123
	JenkinsFileFromType        string `json:"JenkinsFileFromType"`        // STATIC
	JenkinsFilePath            string `json:"JenkinsFilePath"`            // Jenkinsfile
	JenkinsFileStaticContent   string `json:"JenkinsFileStaticContent"`   // 文件内容过长请忽略
	JobFromType                string `json:"JobFromType"`                // TCB_FRAMEWORK
	Name                       string `json:"Name"`                       // hugoo201905135027hello64f84173-321b-4f36-aae7-5b87bfd9b7de
	ProjectID                  int64  `json:"ProjectId"`                  // 645819
	TriggerRemind              string `json:"TriggerRemind"`              // ALWAYS
}

func (req *ModifyCodingCIJobReq) SetAction() string {
	return "ModifyCodingCIJob"
}

type ModifyCodingCIJobResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ModifyCodingCIJob 修改构建计划
func (c *CiClient) ModifyCodingCIJob(req ModifyCodingCIJobReq) (resp ModifyCodingCIJobResp, err error) {
	err = c.Call(&req, &resp)
	return
}
