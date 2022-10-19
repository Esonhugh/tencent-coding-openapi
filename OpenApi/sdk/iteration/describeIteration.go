package iteration

type DescribeIterationReq struct {
	Action        string `json:"Action"`        // DescribeIteration
	IterationCode int64  `json:"IterationCode"` // 1
	ProjectName   string `json:"ProjectName"`   // project-demo
}

func (req *DescribeIterationReq) SetAction() string {
	req.Action = "DescribeIteration"
	return req.Action
}

type DescribeIterationResp struct {
	Response struct {
		Iteration struct {
			Assignee         int64  `json:"Assignee"`         // 0
			Code             int64  `json:"Code"`             // 8
			CompletedCount   int64  `json:"CompletedCount"`   // 0
			CompletedPercent int64  `json:"CompletedPercent"` // 0
			Completer        int64  `json:"Completer"`        // 0
			CreatedAt        int64  `json:"CreatedAt"`        // 1.599027192347e+12
			Creator          int64  `json:"Creator"`          // 1
			Deleter          int64  `json:"Deleter"`          // 0
			EndAt            int64  `json:"EndAt"`            // -2.88e+07
			Goal             string `json:"Goal"`             //
			Name             string `json:"Name"`             // TT
			ProcessingCount  int64  `json:"ProcessingCount"`  // 0
			StartAt          int64  `json:"StartAt"`          // -2.88e+07
			Starter          int64  `json:"Starter"`          // 0
			Status           string `json:"Status"`           // WAIT_PROCESS
			UpdatedAt        int64  `json:"UpdatedAt"`        // 1.599027192347e+12
			WaitProcessCount int64  `json:"WaitProcessCount"` // 0
		} `json:"Iteration"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeIteration 迭代详情
func (i *IterationClient) DescribeIteration(req DescribeIterationReq) (resp DescribeIterationResp, err error) {
	err = i.Call(&req, &resp)
	return
}
