package ci

type TriggerCodingCIBuildReq struct {
	Action   string `json:"Action"`   // TriggerCodingCIBuild
	JobID    int64  `json:"JobId"`    // 23
	Revision string `json:"Revision"` // master
}

func (req *TriggerCodingCIBuildReq) SetAction() string {
	req.Action = "TriggerCodingCIBuild"
	return req.Action
}

type TriggerCodingCIBuildResp struct {
	Response struct {
		Data struct {
			Build struct {
				Branch             string `json:"Branch"`             //
				Cause              string `json:"Cause"`              // coding 手动触发
				CodingCIID         string `json:"CodingCIId"`         // cci-2-605077
				CommitID           string `json:"CommitId"`           // 14c8e6e51ea01fc916dbcaf416f79a68f42c7634
				CreatedAt          int64  `json:"CreatedAt"`          // 1.589084450366e+12
				Duration           int64  `json:"Duration"`           // 0
				FailedMessage      string `json:"FailedMessage"`      //
				ID                 int64  `json:"Id"`                 // 55
				JenkinsFileContent string `json:"JenkinsFileContent"` // pipeline {
				JobID              int64  `json:"JobId"`              // 23
				Number             int64  `json:"Number"`             // 2
				Status             string `json:"Status"`             // QUEUED
				StatusNode         string `json:"StatusNode"`         //
				TestResult         struct {
					Duration   int64 `json:"Duration"`   // 0
					Empty      bool  `json:"Empty"`      // false
					FailCount  int64 `json:"FailCount"`  // 0
					PassCount  int64 `json:"PassCount"`  // 0
					SkipCount  int64 `json:"SkipCount"`  // 0
					TotalCount int64 `json:"TotalCount"` // 0
				} `json:"TestResult"`
			} `json:"Build"`
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// TriggerCodingCIBuild 触发构建
func (c *CiClient) TriggerCodingCIBuild(req TriggerCodingCIBuildReq) (resp TriggerCodingCIBuildResp, err error) {
	err = c.Call(&req, &resp)
	return
}
