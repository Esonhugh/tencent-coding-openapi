package test

type ArchiveTestRunReq struct {
	Action      string `json:"Action"`      // ArchiveTestRun
	ProjectName string `json:"ProjectName"` // project name
	RunID       int64  `json:"RunId"`       // 1
}

func (req *ArchiveTestRunReq) SetAction() string {
	req.Action = "ArchiveTestRun"
	return req.Action
}

type ArchiveTestRunResp struct {
	Response struct {
		Data struct {
			Run struct {
				AssignedToID  int64       `json:"AssignedToId"`  // 0
				BlockedCount  int64       `json:"BlockedCount"`  // 0
				CompletedAt   interface{} `json:"CompletedAt"`   // <nil>
				CreatedAt     string      `json:"CreatedAt"`     // 2020-09-16 15:10:08
				CreatedBy     int64       `json:"CreatedBy"`     // 1
				Description   string      `json:"Description"`   //
				FailedCount   int64       `json:"FailedCount"`   // 0
				ID            int64       `json:"Id"`            // 89
				IncludeAll    bool        `json:"IncludeAll"`    // true
				IsCompleted   bool        `json:"IsCompleted"`   // false
				IterationID   interface{} `json:"IterationId"`   // <nil>
				Name          string      `json:"Name"`          // run demo
				PassedCount   int64       `json:"PassedCount"`   // 0
				RetestCount   int64       `json:"RetestCount"`   // 0
				UntestedCount int64       `json:"UntestedCount"` // 18
			} `json:"Run"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // daba6d8a-0a62-e6fd-f502-a4594888561a
	} `json:"Response"`
}

// ArchiveTestRun 归档测试计划
func (t *TestClient) ArchiveTestRun(req ArchiveTestRunReq) (resp ArchiveTestRunResp, err error) {
	err = t.Call(&req, &resp)
	return
}
