package ci

type CreateCodingCIJobReq struct {
	Action                     string `json:"Action"`                     // CreateCodingCIJob
	AutoCancelSameMergeRequest bool   `json:"AutoCancelSameMergeRequest"` // true
	AutoCancelSameRevision     bool   `json:"AutoCancelSameRevision"`     // true
	BranchRegex                string `json:"BranchRegex"`                //
	BranchSelector             string `json:"BranchSelector"`             //
	DepotID                    int64  `json:"DepotId"`                    // 0
	DepotType                  string `json:"DepotType"`                  // NONE
	ExecuteIn                  string `json:"ExecuteIn"`                  // CVM
	HookType                   string `json:"HookType"`                   // DEFAULT
	JenkinsFileFromType        string `json:"JenkinsFileFromType"`        // STATIC
	JenkinsFilePath            string `json:"JenkinsFilePath"`            // Jenkinsfile
	JenkinsFileStaticContent   string `json:"JenkinsFileStaticContent"`   // 文件内容过长请忽略
	JobFromType                string `json:"JobFromType"`                // CODING
	Name                       string `json:"Name"`                       // hugoo201905135027hello64f84173-321b-4f36-aae7-5b87bfd9b7de
	ProjectID                  int64  `json:"ProjectId"`                  // 645819
	TriggerRemind              string `json:"TriggerRemind"`              // ALWAYS
}

func (req *CreateCodingCIJobReq) SetAction() string {
	req.Action = "CreateCodingCIJob"
	return req.Action
}

type CreateCodingCIJobResp struct {
	Response struct {
		Data struct {
			ID int64 `json:"Id"` // 24
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// CreateCodingCIJob 创建构建计划
func (c *CiClient) CreateCodingCIJob(req CreateCodingCIJobReq) (resp CreateCodingCIJobResp, err error) {
	err = c.Call(&req, &resp)
	return
}
