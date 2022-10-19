package ci

type DescribeCodingCIBuildReq struct {
	Action  string `json:"Action"`  // DescribeCodingCIBuild
	BuildID int64  `json:"BuildId"` // 1
}

func (req *DescribeCodingCIBuildReq) SetAction() string {
	req.Action = "DescribeCodingCIBuild"
	return req.Action
}

type DescribeCodingCIBuildResp struct {
	Response struct {
		Build struct {
			Branch             string `json:"Branch"`             //
			Cause              string `json:"Cause"`              // coding 手动触发
			CodingCIID         string `json:"CodingCIId"`         // cci-2-605077
			CommitID           string `json:"CommitId"`           // 14c8e6e51ea01fc916dbcaf416f79a68f42c7634
			CreatedAt          int64  `json:"CreatedAt"`          // 1.58908445e+12
			Duration           int64  `json:"Duration"`           // 8000
			FailedMessage      string `json:"FailedMessage"`      //
			ID                 int64  `json:"Id"`                 // 55
			JenkinsFileContent string `json:"JenkinsFileContent"` // pipeline {
			JobID              int64  `json:"JobId"`              // 23
			Number             int64  `json:"Number"`             // 2
			Status             string `json:"Status"`             // SUCCEED
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
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeCodingCIBuild 查询构建记录详情
func (c *CiClient) DescribeCodingCIBuild(req DescribeCodingCIBuildReq) (resp DescribeCodingCIBuildResp, err error) {
	err = c.Call(&req, &resp)
	return
}
