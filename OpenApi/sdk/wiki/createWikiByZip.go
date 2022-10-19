package wiki

type CreateWikiByZipReq struct {
	Action        string `json:"Action"`        // CreateWikiByZip
	AuthToken     string `json:"AuthToken"`     // 5b3848e2dfa921e183a071df84aa513d2018031b4df6786dfa7a214a9db855b1
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	FileName      string `json:"FileName"`      // filename.zip
	Key           string `json:"Key"`           // 9fc310d0-3837-11eb-bca5-bd71c937c060.zip
	ParentIid     int64  `json:"ParentIid"`     // 0
	ProjectName   string `json:"ProjectName"`   // demo-project
	Time          int64  `json:"Time"`          // 1.605247760074e+12
}

func (req *CreateWikiByZipReq) SetAction() string {
	req.Action = "CreateWikiByZip"
	return req.Action
}

type CreateWikiByZipResp struct {
	Response struct {
		JobID     string `json:"JobId"`            // 5b3848e2dfa921e183a07
		RequestID int64  `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// CreateWikiByZip 通过 ZIP 包上传
func (w *WikiClient) CreateWikiByZip(req CreateWikiByZipReq) (resp CreateWikiByZipResp, err error) {
	err = w.Call(&req, &resp)
	return
}
