package wiki

type ModifyWikiByZipReq struct {
	Action        string `json:"Action"`        // ModifyWikiByZip
	AuthToken     string `json:"AuthToken"`     // 5b3848e2dfa921e183a071df84aa513d2018031b4df6786dfa7a214a9db855b1
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	FileName      string `json:"FileName"`      // filename.zip
	Iid           int64  `json:"Iid"`           // 11
	Key           string `json:"Key"`           // 9fc310d0-3837-11eb-bca5-bd71c937c060.zip
	ProjectName   string `json:"ProjectName"`   // demo-project
	Time          int64  `json:"Time"`          // 1.605247760074e+12
}

func (req *ModifyWikiByZipReq) SetAction() string {
	req.Action = "ModifyWikiByZip"
	return req.Action
}

type ModifyWikiByZipResp struct {
	Response struct {
		JobID     string `json:"JobId"`            // 5b3848e2dfa921e183a071df84aa
		RequestID int64  `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ModifyWikiByZip 通过 ZIP 包更新 Wiki
func (w *WikiClient) ModifyWikiByZip(req ModifyWikiByZipReq) (resp ModifyWikiByZipResp, err error) {
	err = w.Call(&req, &resp)
	return
}
