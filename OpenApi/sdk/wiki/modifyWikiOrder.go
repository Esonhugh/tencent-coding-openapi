package wiki

type ModifyWikiOrderReq struct {
	Action        string `json:"Action"`        // ModifyWikiOrder
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	Iid           int64  `json:"Iid"`           // 1
	ProjectName   string `json:"ProjectName"`   // demo-project
}

func (req *ModifyWikiOrderReq) SetAction() string {
	req.Action = "ModifyWikiOrder"
	return req.Action
}

type ModifyWikiOrderResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ModifyWikiOrder 修改父级顺序
func (w *WikiClient) ModifyWikiOrder(req ModifyWikiOrderReq) (resp ModifyWikiOrderResp, err error) {
	err = w.Call(&req, &resp)
	return
}
