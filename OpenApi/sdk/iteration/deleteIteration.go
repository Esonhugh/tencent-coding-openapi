package iteration

type DeleteIterationReq struct {
	Action        string `json:"Action"`        // DeleteIteration
	IterationCode int64  `json:"IterationCode"` // 1
	ProjectName   string `json:"ProjectName"`   // demo-project
}

func (req *DeleteIterationReq) SetAction() string {
	req.Action = "DeleteIteration"
	return req.Action
}

type DeleteIterationResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// DeleteIteration 删除迭代
func (i *IterationClient) DeleteIteration(req DeleteIterationReq) (resp DeleteIterationResp, err error) {
	err = i.Call(&req, &resp)
	return
}
