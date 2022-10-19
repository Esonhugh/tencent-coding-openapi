package wiki

type DeleteWikiReq struct {
	Action        string `json:"Action"`        // DeleteWiki
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	Iid           int64  `json:"Iid"`           // 1
	ProjectName   string `json:"ProjectName"`   // demo-project
}

func (req *DeleteWikiReq) SetAction() string {
	req.Action = "DeleteWiki"
	return req.Action
}

type DeleteWikiResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DeleteWiki 删除 Wiki
func (w *WikiClient) DeleteWiki(req DeleteWikiReq) (resp DeleteWikiResp, err error) {
	err = w.Call(&req, &resp)
	return
}
