package test

type DescribeTestRunListReq struct {
	Action      string `json:"Action"`      // DescribeTestRunList
	ProjectName string `json:"ProjectName"` // project name
}

func (req *DescribeTestRunListReq) SetAction() string {
	req.Action = "DescribeTestRunList"
	return req.Action
}

type DescribeTestRunListResp struct {
	Response struct {
		Data struct {
			Runs []struct {
				AssignedToID        int64       `json:"AssignedToId"`          // 0
				BlockedCount        int64       `json:"BlockedCount"`          // 0
				CompletedAt         interface{} `json:"CompletedAt"`           // <nil>
				ConfigEnvironmentID int64       `json:"ConfigEnvironmentId"`   // 2
				CreatedAt           string      `json:"CreatedAt"`             // 2021-06-28 11:44:02
				CreatedBy           int64       `json:"CreatedBy"`             // 1
				Days                int64       `json:"Days"`                  // 9
				Description         string      `json:"Description"`           // 描述
				ExecuteType         int64       `json:"ExecuteType"`           // 1
				FailedCount         int64       `json:"FailedCount"`           // 0
				GitDepotID          int64       `json:"GitDepotId"`            // 1
				GitDepotName        string      `json:"GitDepotName"`          // git-test
				GitReleaseID        int64       `json:"GitReleaseId"`          // 1580
				GitReleaseName      float64     `json:"GitReleaseName,string"` // 20210621.1
				GitReleaseState     int64       `json:"GitReleaseState"`       // 0
				ID                  int64       `json:"Id"`                    // 1747
				IncludeAll          bool        `json:"IncludeAll"`            // false
				IsCompleted         bool        `json:"IsCompleted"`           // false
				IterationID         interface{} `json:"IterationId"`           // <nil>
				IterationName       interface{} `json:"IterationName"`         // <nil>
				Name                string      `json:"Name"`                  // 有个计划
				PassedCount         int64       `json:"PassedCount"`           // 0
				RetestCount         int64       `json:"RetestCount"`           // 0
				SectionID           int64       `json:"SectionId"`             // -1
				SectionName         interface{} `json:"SectionName"`           // <nil>
				State               int64       `json:"State"`                 // 0
				UntestedCount       int64       `json:"UntestedCount"`         // 2
			} `json:"Runs"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 5ad95c11-a31e-db3e-76ca-59a7bce92958
	} `json:"Response"`
}

// DescribeTestRunList 测试计划列表
func (t *TestClient) DescribeTestRunList(req DescribeTestRunListReq) (resp DescribeTestRunListResp, err error) {
	err = t.Call(&req, &resp)
	return
}
