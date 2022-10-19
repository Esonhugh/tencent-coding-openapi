package issue

type DescribeIssueReq struct {
	Action      string `json:"Action"`      // DescribeIssue
	IssueCode   int64  `json:"IssueCode"`   // 1
	ProjectName string `json:"ProjectName"` // demo-project
}

func (req *DescribeIssueReq) SetAction() string {
	req.Action = "DescribeIssue"
	return req.Action
}

type DescribeIssueResp struct {
	Response struct {
		Issue struct {
			Assignee struct {
				Avatar    string `json:"Avatar"`    //
				Email     string `json:"Email"`     //
				GlobalKey string `json:"GlobalKey"` //
				ID        int64  `json:"Id"`        // 0
				Media     string `json:"Media"`     //
				Name      string `json:"Name"`      //
				Status    int64  `json:"Status"`    // 0
			} `json:"Assignee"`
			Code        int64 `json:"Code"`        // 1
			CompletedAt int64 `json:"CompletedAt"` // 0
			CreatedAt   int64 `json:"CreatedAt"`   // 1.597384413e+12
			Creator     struct {
				Avatar    string `json:"Avatar"`    // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Email     string `json:"Email"`     //
				GlobalKey string `json:"GlobalKey"` // coding
				ID        int64  `json:"Id"`        // 1
				Media     string `json:"Media"`     // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Name      string `json:"Name"`      // coding
				Status    int64  `json:"Status"`    // 2
			} `json:"Creator"`
			CustomFields []struct {
				ID          int64  `json:"Id"`          // 16
				ValueString string `json:"ValueString"` // [{"Id":1,"Status":2,"GlobalKey":"coding","Avatar":"https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0","Media":"https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0","Name":"coding","Email":"coding@coding.com","Name_pinyin":""}]
			} `json:"CustomFields"`
			DefectType struct {
				IconURL string `json:"IconUrl"` //
				ID      int64  `json:"Id"`      // 0
				Name    string `json:"Name"`    //
			} `json:"DefectType"`
			Description string `json:"Description"` // # 标题1
			DueDate     int64  `json:"DueDate"`     // 1.596643199e+12
			Epic        struct {
				Assignee struct {
					Avatar    string `json:"Avatar"`    //
					Email     string `json:"Email"`     //
					GlobalKey string `json:"GlobalKey"` //
					ID        int64  `json:"Id"`        // 0
					Media     string `json:"Media"`     //
					Name      string `json:"Name"`      //
					Status    int64  `json:"Status"`    // 0
				} `json:"Assignee"`
				Code            int64  `json:"Code"`            // 0
				IssueStatusID   int64  `json:"IssueStatusId"`   // 0
				IssueStatusName string `json:"IssueStatusName"` //
				Name            string `json:"Name"`            //
				Priority        string `json:"Priority"`        //
				Type            string `json:"Type"`            //
			} `json:"Epic"`
			Files           []interface{} `json:"Files"`           // <nil>
			IssueStatusID   int64         `json:"IssueStatusId"`   // 4
			IssueStatusName string        `json:"IssueStatusName"` // 未开始
			IssueStatusType string        `json:"IssueStatusType"` // TODO
			IssueTypeDetail struct {
				Description string `json:"Description"` //
				ID          int64  `json:"Id"`          // 1
				IsSystem    bool   `json:"IsSystem"`    // false
				IssueType   string `json:"IssueType"`   // REQUIREMENT
				Name        string `json:"Name"`        // 用户故事
			} `json:"IssueTypeDetail"`
			Iteration struct {
				Code   int64  `json:"Code"`   // 1
				Name   string `json:"Name"`   // Q1 Spring1
				Status string `json:"Status"` // PROCESSING
			} `json:"Iteration"`
			IterationID int64         `json:"IterationId"` // 1
			Labels      []interface{} `json:"Labels"`      // <nil>
			Name        string        `json:"Name"`        // 测试需求1
			Parent      struct {
				Assignee struct {
					Avatar    string `json:"Avatar"`    //
					Email     string `json:"Email"`     //
					GlobalKey string `json:"GlobalKey"` //
					ID        int64  `json:"Id"`        // 0
					Media     string `json:"Media"`     //
					Name      string `json:"Name"`      //
					Status    int64  `json:"Status"`    // 0
				} `json:"Assignee"`
				Code            int64  `json:"Code"`            // 0
				IssueStatusID   int64  `json:"IssueStatusId"`   // 0
				IssueStatusName string `json:"IssueStatusName"` //
				Name            string `json:"Name"`            //
				Priority        string `json:"Priority"`        //
				Type            string `json:"Type"`            //
			} `json:"Parent"`
			ParentType    string `json:"ParentType"`      // REQUIREMENT
			Priority      int64  `json:"Priority,string"` // 0
			ProjectModule struct {
				ID   int64  `json:"Id"`   // 0
				Name string `json:"Name"` //
			} `json:"ProjectModule"`
			RequirementType struct {
				ID   int64  `json:"Id"`   // 0
				Name string `json:"Name"` //
			} `json:"RequirementType"`
			StartDate  int64         `json:"StartDate"`  // 0
			StoryPoint string        `json:"StoryPoint"` //
			SubTasks   []interface{} `json:"SubTasks"`   // <nil>
			ThirdLinks []struct {
				ID        int64  `json:"Id"`        // 1
				Link      string `json:"Link"`      // https://modao.cc/app/7fb28d9af13827fd009f401579cbdc3cef2a456a/embed/v2
				ThirdType string `json:"ThirdType"` // MODAO
				Title     string `json:"Title"`     // 墨刀 1
			} `json:"ThirdLinks"`
			Type         string        `json:"Type"`         // REQUIREMENT
			UpdatedAt    int64         `json:"UpdatedAt"`    // 1.599035298e+12
			Watchers     []interface{} `json:"Watchers"`     // <nil>
			WorkingHours int64         `json:"WorkingHours"` // 0
		} `json:"Issue"`
		RequestID string `json:"RequestId"` // d2e3c64c-99f8-1ef2-c148-36f7d6a81cfe
	} `json:"Response"`
}

// DescribeIssue 查询事项详情
func (i *IssueClient) DescribeIssue(req DescribeIssueReq) (resp DescribeIssueResp, err error) {
	err = i.Call(&req, &resp)
	return
}
