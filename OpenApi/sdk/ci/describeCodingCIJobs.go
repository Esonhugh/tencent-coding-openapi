package ci

type DescribeCodingCIJobsReq struct {
	Action string `json:"Action"` // DescribeCodingCIJobs
	Filter []struct {
		Name  string   `json:"Name"`  // Name
		Value []string `json:"Value"` // job-create
	} `json:"Filter"`
	ProjectID int64 `json:"ProjectId"` // 123
}

func (req *DescribeCodingCIJobsReq) SetAction() string {
	req.Action = "DescribeCodingCIJobs"
	return req.Action
}

type DescribeCodingCIJobsResp struct {
	Response struct {
		JobList []struct {
			AlwaysUserIDList           []interface{} `json:"AlwaysUserIdList"`           // <nil>
			AutoCancelSameMergeRequest bool          `json:"AutoCancelSameMergeRequest"` // true
			AutoCancelSameRevision     bool          `json:"AutoCancelSameRevision"`     // true
			BranchRegex                string        `json:"BranchRegex"`                // ^refs/heads/master$
			BranchSelector             string        `json:"BranchSelector"`             // master
			BuildFailUserIDList        []interface{} `json:"BuildFailUserIdList"`        // <nil>
			CachePathList              []interface{} `json:"CachePathList"`              // <nil>
			CreatedAt                  int64         `json:"CreatedAt"`                  // 0
			CreatorID                  int64         `json:"CreatorId"`                  // 1
			DepotHttpsURL              string        `json:"DepotHttpsUrl"`              // http://e.coding.9.134.115.58.nip.io/codingcorp/test-1.git
			DepotID                    int64         `json:"DepotId"`                    // 1
			DepotName                  string        `json:"DepotName"`                  // test-1
			DepotSshURL                string        `json:"DepotSshUrl"`                // git@e.coding.9.134.115.58.nip.io:codingcorp/test-1.git
			DepotType                  string        `json:"DepotType"`                  // CODING
			DepotWebURL                string        `json:"DepotWebUrl"`                // http://codingcorp.coding.9.134.115.58.nip.io/p/test-1/d/test-1/git
			DockerBuildPath            string        `json:"DockerBuildPath"`            //
			DockerBuildTag             string        `json:"DockerBuildTag"`             //
			DockerFilePath             string        `json:"DockerFilePath"`             //
			EnvList                    []struct {
				Name      string `json:"Name"`         // a
				Sensitive bool   `json:"Sensitive"`    // false
				Value     int64  `json:"Value,string"` // 1
			} `json:"EnvList"`
			ExecuteIn                string        `json:"ExecuteIn"`                // CVM
			GroupName                string        `json:"GroupName"`                // test1
			HookType                 string        `json:"HookType"`                 // DEFAULT
			ID                       int64         `json:"Id"`                       // 23
			JenkinsFileFromType      string        `json:"JenkinsFileFromType"`      // STATIC
			JenkinsFilePath          string        `json:"JenkinsFilePath"`          // Jenkinsfile
			JenkinsFileStaticContent string        `json:"JenkinsFileStaticContent"` // pipeline {
			JobFromType              string        `json:"JobFromType"`              // CODING
			Name                     string        `json:"Name"`                     // job-create
			ProjectID                int64         `json:"ProjectId"`                // 450
			ProjectName              string        `json:"ProjectName"`              // test-1
			ScheduleList             []interface{} `json:"ScheduleList"`             // <nil>
			TriggerMethodList        []string      `json:"TriggerMethodList"`        // MR_CHANGE
			TriggerRemind            string        `json:"TriggerRemind"`            // ALWAYS
			UpdatedAt                int64         `json:"UpdatedAt"`                // 0
		} `json:"JobList"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeCodingCIJobs 查询单个项目下的所有构建计划
func (c *CiClient) DescribeCodingCIJobs(req DescribeCodingCIJobsReq) (resp DescribeCodingCIJobsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
