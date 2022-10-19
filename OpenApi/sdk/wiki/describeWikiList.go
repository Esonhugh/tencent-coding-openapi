package wiki

type DescribeWikiListReq struct {
	Action        string `json:"Action"`        // DescribeWikiList
	Authorization string `json:"Authorization"` // 8c044154bd5c7cadb5720944c733b018d0984b3
	ProjectName   string `json:"ProjectName"`   // demo-project
}

func (req *DescribeWikiListReq) SetAction() string {
	req.Action = "DescribeWikiList"
	return req.Action
}

type DescribeWikiListResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
		Data      []struct {
			ID          int64 `json:"id"`          // 359411
			CanMaintain bool  `json:"canMaintain"` // false
			CanRead     bool  `json:"canRead"`     // false
			Children    []struct {
				ID          int64 `json:"id"`          // 359413
				CanMaintain bool  `json:"canMaintain"` // true
				CanRead     bool  `json:"canRead"`     // false
				Children    []struct {
					ID          int64         `json:"id"`          // 359414
					CanMaintain bool          `json:"canMaintain"` // true
					CanRead     bool          `json:"canRead"`     // false
					Children    []interface{} `json:"children"`    // <nil>
					CreatedAt   int64         `json:"createdAt"`   // 1.578036875e+12
					Creator     struct {
						ID        int64  `json:"id"`        // 38981
						Avatar    string `json:"avatar"`    // https://dn-coding-net-production-static.codehub.cn/d466b293-5637-48e7-b485-e979d125d7a6.jpg?imageMogr2/auto-orient/format/jpeg/crop/!639x639a0a0
						Email     string `json:"email"`     // wangjiajun@coding.net
						GlobalKey string `json:"globalKey"` // wangjiajun
						Name      string `json:"name"`      // 王佳骏
						Status    int64  `json:"status"`    // 1
					} `json:"creator"`
					Editor struct {
						ID        int64  `json:"id"`        // 39018
						Avatar    string `json:"avatar"`    // https://dn-coding-net-production-static.codehub.cn/86e98800-d835-41d8-8142-b7fca0550bf7.jpg?imageMogr2/auto-orient/format/jpeg/crop/!410x410a0a0
						Email     string `json:"email"`     // caiyuezhang@coding.net
						GlobalKey string `json:"globalKey"` // caiyuezhang
						Name      string `json:"name"`      // 蔡悦彰
						Status    int64  `json:"status"`    // 1
					} `json:"editor"`
					Iid          int64  `json:"iid"`          // 26
					IsShared     bool   `json:"isShared"`     // false
					IsTreeShared bool   `json:"isTreeShared"` // false
					Order        int64  `json:"order"`        // 1
					ParentIid    int64  `json:"parentIid"`    // 25
					Title        string `json:"title"`        // Wiki 产品
					UpdatedAt    int64  `json:"updatedAt"`    // 1.586787643e+12
					VisibleRange string `json:"visibleRange"` // INHERIT
				} `json:"children"`
				CreatedAt int64 `json:"createdAt"` // 1.578036875e+12
				Creator   struct {
					ID        int64  `json:"id"`        // 38981
					Avatar    string `json:"avatar"`    // https://dn-coding-net-production-static.codehub.cn/d466b293-5637-48e7-b485-e979d125d7a6.jpg?imageMogr2/auto-orient/format/jpeg/crop/!639x639a0a0
					Email     string `json:"email"`     // wangjiajun@coding.net
					GlobalKey string `json:"globalKey"` // wangjiajun
					Name      string `json:"name"`      // 王佳骏
					Status    int64  `json:"status"`    // 1
				} `json:"creator"`
				Editor struct {
					ID        int64  `json:"id"`        // 39018
					Avatar    string `json:"avatar"`    // https://dn-coding-net-production-static.codehub.cn/86e98800-d835-41d8-8142-b7fca0550bf7.jpg?imageMogr2/auto-orient/format/jpeg/crop/!410x410a0a0
					Email     string `json:"email"`     // caiyuezhang@coding.net
					GlobalKey string `json:"globalKey"` // caiyuezhang
					Name      string `json:"name"`      // 蔡悦彰
					Status    int64  `json:"status"`    // 1
				} `json:"editor"`
				Iid          int64  `json:"iid"`          // 25
				IsShared     bool   `json:"isShared"`     // false
				IsTreeShared bool   `json:"isTreeShared"` // false
				Order        int64  `json:"order"`        // 1
				ParentIid    int64  `json:"parentIid"`    // 24
				Title        string `json:"title"`        // 产品文档
				UpdatedAt    int64  `json:"updatedAt"`    // 1.586787643e+12
				VisibleRange string `json:"visibleRange"` // READ_ONLY
			} `json:"children"`
			Iid          int64       `json:"iid"`          // 24
			IsShared     bool        `json:"isShared"`     // false
			IsTreeShared bool        `json:"isTreeShared"` // true
			Order        int64       `json:"order"`        // 1
			ParentIid    int64       `json:"parentIid"`    // 0
			Title        interface{} `json:"title"`        // <nil>
			VisibleRange string      `json:"visibleRange"` // HIDDEN
		} `json:"data"`
	} `json:"Response"`
}

// DescribeWikiList 获取列表详情
func (w *WikiClient) DescribeWikiList(req DescribeWikiListReq) (resp DescribeWikiListResp, err error) {
	err = w.Call(&req, &resp)
	return
}
