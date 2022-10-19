package wiki

type DescribeUpdateJobStatusReq struct {
	Action        string `json:"Action"`        // DescribeUpdateJobStatus
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	JobID         string `json:"JobId"`         // asdaqweqsdaddsdsadsa
	ProjectName   string `json:"ProjectName"`   // demo-project
}

func (req *DescribeUpdateJobStatusReq) SetAction() string {
	req.Action = "DescribeUpdateJobStatus"
	return req.Action
}

type DescribeUpdateJobStatusResp struct {
	Response struct {
		JobID     string `json:"JobId"`            // 5b3848e2dfa921e183a071df84aa
		RequestID int64  `json:"RequestId,string"` // 1
		Status    string `json:"Status"`           // WAIT_PROCESS
	} `json:"Response"`
}

// DescribeUpdateJobStatus 通过 ZIP 包更新任务状态查询
func (w *WikiClient) DescribeUpdateJobStatus(req DescribeUpdateJobStatusReq) (resp DescribeUpdateJobStatusResp, err error) {
	err = w.Call(&req, &resp)
	return
}
