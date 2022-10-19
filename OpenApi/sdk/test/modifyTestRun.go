package test

type ModifyTestRunReq struct {
	Action      string  `json:"Action"`      // ModifyTestRun
	Cases       []int64 `json:"Cases"`       // 1
	IncludeAll  bool    `json:"IncludeAll"`  // false
	Name        string  `json:"Name"`        // run demo
	ProjectName string  `json:"ProjectName"` // project name
	RunID       int64   `json:"RunId"`       // 1
}

func (req *ModifyTestRunReq) SetAction() string {
	req.Action = "ModifyTestRun"
	return req.Action
}

type ModifyTestRunResp struct {
	Response struct {
		Data struct {
			Run struct {
				AssignedToID        int64       `json:"AssignedToId"`          // 0
				BlockedCount        int64       `json:"BlockedCount"`          // 0
				CompletedAt         interface{} `json:"CompletedAt"`           // <nil>
				ConfigEnvironmentID int64       `json:"ConfigEnvironmentId"`   // 2
				CreatedAt           string      `json:"CreatedAt"`             // 2020-09-16 15:10:08
				CreatedBy           int64       `json:"CreatedBy"`             // 1
				Days                int64       `json:"Days"`                  // 0
				Description         string      `json:"Description"`           //
				ExecuteType         int64       `json:"ExecuteType"`           // 1
				FailedCount         int64       `json:"FailedCount"`           // 0
				GitDepotID          int64       `json:"GitDepotId"`            // 1
				GitDepotName        string      `json:"GitDepotName"`          // git-test
				GitReleaseID        int64       `json:"GitReleaseId"`          // 1580
				GitReleaseName      float64     `json:"GitReleaseName,string"` // 20210621.1
				GitReleaseState     int64       `json:"GitReleaseState"`       // 0
				ID                  int64       `json:"Id"`                    // 89
				IncludeAll          bool        `json:"IncludeAll"`            // true
				IsCompleted         bool        `json:"IsCompleted"`           // false
				IterationID         interface{} `json:"IterationId"`           // <nil>
				IterationName       interface{} `json:"IterationName"`         // <nil>
				Name                string      `json:"Name"`                  // run demo
				PassedCount         int64       `json:"PassedCount"`           // 0
				RetestCount         int64       `json:"RetestCount"`           // 0
				SectionID           int64       `json:"SectionId"`             // -1
				SectionName         interface{} `json:"SectionName"`           // <nil>
				State               interface{} `json:"State"`                 // <nil>
				UntestedCount       int64       `json:"UntestedCount"`         // 18
			} `json:"Run"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // daba6d8a-0a62-e6fd-f502-a4594888561a
	} `json:"Response"`
}

// ModifyTestRun 修改测试计划
func (t *TestClient) ModifyTestRun(req ModifyTestRunReq) (resp ModifyTestRunResp, err error) {
	err = t.Call(&req, &resp)
	return
}
