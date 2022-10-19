package wiki

type DescribeWikiReq struct {
	Action      string `json:"Action"`      // DescribeWiki
	Iid         int64  `json:"Iid"`         // 7
	ProjectName string `json:"ProjectName"` // demo-project
	VersionID   int64  `json:"VersionId"`   // 1
}

func (req *DescribeWikiReq) SetAction() string {
	req.Action = "DescribeWiki"
	return req.Action
}

type DescribeWikiResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
		Data      struct {
			ID        int64  `json:"id"`        // 10
			Content   string `json:"content"`   //
			CreatedAt int64  `json:"createdAt"` // 1.578378234037e+12
			Creator   struct {
				ID        int64  `json:"id"`         // 1
				Avatar    string `json:"avatar"`     // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Email     string `json:"email"`      // 1111@coding.com
				GlobalKey string `json:"global_key"` // coding
				Name      string `json:"name"`       // coding
				Phone     int64  `json:"phone"`      // 0
				Status    int64  `json:"status"`     // 1
				TeamID    int64  `json:"teamId"`     // 0
			} `json:"creator"`
			CurrentUserRoleID int64 `json:"currentUserRoleId"` // 100
			CurrentVersion    int64 `json:"currentVersion"`    // 1
			Editor            struct {
				ID        int64  `json:"id"`         // 1
				Avatar    string `json:"avatar"`     // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Email     string `json:"email"`      // 1111@coding.com
				GlobalKey string `json:"global_key"` // coding
				Name      string `json:"name"`       // coding
				Phone     int64  `json:"phone"`      // 0
				Status    int64  `json:"status"`     // 1
				TeamID    int64  `json:"teamId"`     // 0
			} `json:"editor"`
			HistoriesCount int64  `json:"historiesCount"` // 1
			HistoryID      int64  `json:"historyId"`      // 13
			Html           string `json:"html"`           //
			Iid            int64  `json:"iid"`            // 10
			LastVersion    int64  `json:"lastVersion"`    // 1
			Msg            string `json:"msg"`            //
			Order          int64  `json:"order"`          // 3
			ParentIid      int64  `json:"parentIid"`      // 7
			Path           string `json:"path"`           // 1/2/7/10
			Reminded       bool   `json:"reminded"`       // false
			Title          string `json:"title"`          // 不卡了
			UpdatedAt      int64  `json:"updatedAt"`      // 1.578378234041e+12
		} `json:"data"`
	} `json:"Response"`
}

// DescribeWiki 获取详情
func (w *WikiClient) DescribeWiki(req DescribeWikiReq) (resp DescribeWikiResp, err error) {
	err = w.Call(&req, &resp)
	return
}
