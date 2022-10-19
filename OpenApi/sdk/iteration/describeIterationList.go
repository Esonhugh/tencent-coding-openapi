package iteration

type DescribeIterationListReq struct {
	Action      string   `json:"Action"`      // DescribeIterationList
	ProjectName string   `json:"ProjectName"` // demo-project
	Status      []string `json:"Status"`      // WAIT_PROCESS
}

func (req *DescribeIterationListReq) SetAction() string {
	req.Action = "DescribeIterationList"
	return req.Action
}

type DescribeIterationListResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
		Data      struct {
			List []struct {
				Assignee         int64  `json:"Assignee"`         // 0
				Code             int64  `json:"Code"`             // 8
				CompletedCount   int64  `json:"CompletedCount"`   // 0
				CompletedPercent int64  `json:"CompletedPercent"` // 0
				Completer        int64  `json:"Completer"`        // 0
				CreatedAt        int64  `json:"CreatedAt"`        // 1.599027192e+12
				Creator          int64  `json:"Creator"`          // 1
				Deleter          int64  `json:"Deleter"`          // 0
				EndAt            int64  `json:"EndAt"`            // -2.88e+07
				Goal             string `json:"Goal"`             //
				Name             string `json:"Name"`             // demo iteration
				ProcessingCount  int64  `json:"ProcessingCount"`  // 0
				StartAt          int64  `json:"StartAt"`          // -2.88e+07
				Starter          int64  `json:"Starter"`          // 0
				Status           string `json:"Status"`           // WAIT_PROCESS
				UpdatedAt        int64  `json:"UpdatedAt"`        // 1.599027192e+12
				WaitProcessCount int64  `json:"WaitProcessCount"` // 0
			} `json:"List"`
			Page      int64 `json:"Page"`      // 1
			PageSize  int64 `json:"PageSize"`  // 20
			TotalPage int64 `json:"TotalPage"` // 1
			TotalRow  int64 `json:"TotalRow"`  // 1
		} `json:"data"`
	} `json:"Response"`
}

// DescribeIterationList 迭代列表
func (i *IterationClient) DescribeIterationList(req DescribeIterationListReq) (resp DescribeIterationListResp, err error) {
	err = i.Call(&req, &resp)
	return
}
