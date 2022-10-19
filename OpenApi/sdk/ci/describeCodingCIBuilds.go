package ci

type DescribeCodingCIBuildsReq struct {
	Action     string `json:"Action"`     // DescribeCodingCIBuilds
	JobID      int64  `json:"JobId"`      // 1
	PageNumber int64  `json:"PageNumber"` // 1
	PageSize   int64  `json:"PageSize"`   // 10
}

func (req *DescribeCodingCIBuildsReq) SetAction() string {
	req.Action = "DescribeCodingCIBuilds"
	return req.Action
}

type DescribeCodingCIBuildsResp struct {
	Response struct {
		Data struct {
			BuildList []struct {
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
			} `json:"BuildList"`
			PageNumber int64 `json:"PageNumber"` // 1
			PageSize   int64 `json:"PageSize"`   // 2
			TotalCount int64 `json:"TotalCount"` // 2
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeCodingCIBuilds 获取构建计划的构建列表
func (c *CiClient) DescribeCodingCIBuilds(req DescribeCodingCIBuildsReq) (resp DescribeCodingCIBuildsResp, err error) {
	err = c.Call(&req, &resp)
	return
}
