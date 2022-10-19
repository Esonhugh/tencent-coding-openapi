package test

type DescribeTestRunReq struct {
	Action      string `json:"Action"`      // DescribeTestRun
	ProjectName string `json:"ProjectName"` // project name
	RunID       int64  `json:"RunId"`       // 1
}

func (req *DescribeTestRunReq) SetAction() string {
	req.Action = "DescribeTestRun"
	return req.Action
}

type DescribeTestRunResp struct {
	Response struct {
		Data struct {
			Run struct {
				AssignedToID        int64       `json:"AssignedToId"`          // 0
				BlockedCount        int64       `json:"BlockedCount"`          // 0
				CompletedAt         interface{} `json:"CompletedAt"`           // <nil>
				ConfigEnvironmentID int64       `json:"ConfigEnvironmentId"`   // 1
				CreatedAt           string      `json:"CreatedAt"`             // 2021-06-17 17:32:52
				CreatedBy           int64       `json:"CreatedBy"`             // 1
				Days                int64       `json:"Days"`                  // 19
				Description         string      `json:"Description"`           //
				ExecuteType         int64       `json:"ExecuteType"`           // 1
				FailedCount         int64       `json:"FailedCount"`           // 0
				GitDepotID          int64       `json:"GitDepotId"`            // 1
				GitDepotName        string      `json:"GitDepotName"`          // git-test
				GitReleaseID        int64       `json:"GitReleaseId"`          // 1549
				GitReleaseName      float64     `json:"GitReleaseName,string"` // 20210617.1
				GitReleaseState     int64       `json:"GitReleaseState"`       // 0
				ID                  int64       `json:"Id"`                    // 1550
				IncludeAll          bool        `json:"IncludeAll"`            // false
				IsCompleted         bool        `json:"IsCompleted"`           // false
				IterationID         interface{} `json:"IterationId"`           // <nil>
				IterationName       interface{} `json:"IterationName"`         // <nil>
				Name                string      `json:"Name"`                  // 测试计划
				PassedCount         int64       `json:"PassedCount"`           // 0
				RetestCount         int64       `json:"RetestCount"`           // 0
				SectionID           int64       `json:"SectionId"`             // -1
				SectionName         interface{} `json:"SectionName"`           // <nil>
				State               int64       `json:"State"`                 // 0
				UntestedCount       int64       `json:"UntestedCount"`         // 2
			} `json:"Run"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // daba6d8a-0a62-e6fd-f502-a4594888561a
	} `json:"Response"`
}

// DescribeTestRun 测试计划详情
func (t *TestClient) DescribeTestRun(req DescribeTestRunReq) (resp DescribeTestRunResp, err error) {
	err = t.Call(&req, &resp)
	return
}
