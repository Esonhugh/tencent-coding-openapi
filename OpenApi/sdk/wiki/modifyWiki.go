package wiki

type ModifyWikiReq struct {
	Action        string `json:"Action"`        // ModifyWiki
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	Content       string `json:"Content"`       // 内容
	Iid           int64  `json:"Iid"`           // 1
	Msg           string `json:"Msg"`           // 提交说明
	ProjectName   string `json:"ProjectName"`   // demo-project
	Title         string `json:"Title"`         // Title
}

func (req *ModifyWikiReq) SetAction() string {
	req.Action = "ModifyWiki"
	return req.Action
}

type ModifyWikiResp struct {
	Response struct {
		Data struct {
			CanMaintain bool   `json:"CanMaintain"` // true
			CanRead     bool   `json:"CanRead"`     // true
			Content     string `json:"Content"`     // WIKI
			CreatedAt   string `json:"CreatedAt"`   // 2020-09-22 00:00:00
			Creator     struct {
				Avatar        string `json:"Avatar"`        // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Email         string `json:"Email"`         // coding.com
				GlobalKey     string `json:"GlobalKey"`     // coding
				ID            int64  `json:"Id"`            // 1
				Name          string `json:"Name"`          // coding
				Phone         int64  `json:"Phone,string"`  // 133131313
				Status        int64  `json:"Status"`        // 1
				TeamGlobalKey string `json:"TeamGlobalKey"` // coding
				TeamID        int64  `json:"TeamId"`        // 1
			} `json:"Creator"`
			CurrentVersion int64 `json:"CurrentVersion"` // 1
			Editor         struct {
				Avatar        string `json:"Avatar"`        // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Email         string `json:"Email"`         // coding.com
				GlobalKey     string `json:"GlobalKey"`     // coding
				ID            int64  `json:"Id"`            // 1
				Name          string `json:"Name"`          // coding
				Phone         int64  `json:"Phone,string"`  // 133131313
				Status        int64  `json:"Status"`        // 1
				TeamGlobalKey string `json:"TeamGlobalKey"` // coding
				TeamID        int64  `json:"TeamId"`        // 1
			} `json:"Editor"`
			HistoriesCount     int64  `json:"HistoriesCount"`     // 1
			HistoryID          int64  `json:"HistoryId"`          // 13
			Html               string `json:"Html"`               // WIKI
			ID                 int64  `json:"Id"`                 // 10
			Iid                int64  `json:"Iid"`                // 10
			LastVersion        int64  `json:"LastVersion"`        // 1
			Msg                string `json:"Msg"`                // TEST
			Order              int64  `json:"Order"`              // 0
			ParentIid          int64  `json:"ParentIid"`          // 7
			ParentShared       bool   `json:"ParentShared"`       // true
			ParentVisibleRange string `json:"ParentVisibleRange"` // PUBLIC
			Path               int64  `json:"Path,string"`        // 2091
			Title              string `json:"Title"`              // WIKI-TEST
			UpdatedAt          string `json:"UpdatedAt"`          // 2020-09-22 00:00:00
			VisibleRange       string `json:"VisibleRange"`       // INHERIT
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1111111111
	} `json:"Response"`
}

// ModifyWiki 更新 Wiki
func (w *WikiClient) ModifyWiki(req ModifyWikiReq) (resp ModifyWikiResp, err error) {
	err = w.Call(&req, &resp)
	return
}
