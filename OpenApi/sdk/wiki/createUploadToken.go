package wiki

type CreateUploadTokenReq struct {
	Action        string `json:"Action"`        // CreateUploadToken
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	FileName      string `json:"FileName"`      // filename.zip
	ProjectName   string `json:"ProjectName"`   // demo-project
	Time          int64  `json:"Time"`          // 1.605247760074e+12
}

func (req *CreateUploadTokenReq) SetAction() string {
	req.Action = "CreateUploadToken"
	return req.Action
}

type CreateUploadTokenResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
		Token     struct {
			AuthToken  string `json:"authToken"`   // 5b3848e2dfa921e183a071df84aa513d2018031b4df6786dfa7a214a9db855b1
			Provider   string `json:"provider"`    // TENCENT
			SecretID   string `json:"secretId"`    // AKIDOeYe7WatwJ5Hr4-wIW25kxH007lVmjK
			SecretKey  string `json:"secretKey"`   // ggBkOApOtx8p0cw+dhfo8nU31oj7xX39gkpl
			Time       int64  `json:"time,string"` // 1605247760074
			UpToken    string `json:"upToken"`     // bBoXiOuPiCeuXaeiyxAOK2qfsbSJJv5a393b958ff8094b
			UploadLink string `json:"uploadLink"`  // https://coding-net-production-file-1257242599.cos.ap-shanghai.myqcloud.com
		} `json:"Token"`
	} `json:"Response"`
}

// CreateUploadToken 获取上传文件的 Token
func (w *WikiClient) CreateUploadToken(req CreateUploadTokenReq) (resp CreateUploadTokenResp, err error) {
	err = w.Call(&req, &resp)
	return
}
