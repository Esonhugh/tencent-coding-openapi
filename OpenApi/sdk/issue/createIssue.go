package issue

type CreateIssueReq struct {
	Action            string `json:"Action"` // CreateIssue
	CustomFieldValues []struct {
		Content string `json:"Content"` // 1（普通字段）
		ID      int64  `json:"Id"`      // 3.4929409e+07
	} `json:"CustomFieldValues"`
	Name        string `json:"Name"`            // 一个需求
	Priority    int64  `json:"Priority,string"` // 0
	ProjectName string `json:"ProjectName"`     // demo-project
	Type        string `json:"Type"`            // REQUIREMENT
}

func (req *CreateIssueReq) SetAction() string {
	req.Action = "CreateIssue"
	return req.Action
}

type CreateIssueResp struct {
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
			Code        int64 `json:"Code"`        // 27
			CompletedAt int64 `json:"CompletedAt"` // 0
			CreatedAt   int64 `json:"CreatedAt"`   // 1.598843625826e+12
			Creator     struct {
				Avatar    string `json:"Avatar"`    // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Email     string `json:"Email"`     //
				GlobalKey string `json:"GlobalKey"` // coding
				ID        int64  `json:"Id"`        // 1
				Media     string `json:"Media"`     // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Name      string `json:"Name"`      // coding
				Status    int64  `json:"Status"`    // 2
			} `json:"Creator"`
			CustomFields []interface{} `json:"CustomFields"` // <nil>
			DefectType   struct {
				IconURL string `json:"IconUrl"` //
				ID      int64  `json:"Id"`      // 0
				Name    string `json:"Name"`    //
			} `json:"DefectType"`
			Description string `json:"Description"` //
			DueDate     int64  `json:"DueDate"`     // 0
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
			IterationID int64         `json:"IterationId"` // 0
			Labels      []interface{} `json:"Labels"`      // <nil>
			Name        string        `json:"Name"`        // test-20200826-1
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
			ThirdLinks []interface{} `json:"ThirdLinks"` // <nil>
			Type       string        `json:"Type"`       // REQUIREMENT
			UpdatedAt  int64         `json:"UpdatedAt"`  // 1.598843625826e+12
			Watchers   []struct {
				Avatar    string `json:"Avatar"`    // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Email     string `json:"Email"`     //
				GlobalKey string `json:"GlobalKey"` // coding
				ID        int64  `json:"Id"`        // 1
				Media     string `json:"Media"`     // https://dn-coding-net-production-static.codehub.cn/512b2a62-956b-4ef8-8e84-b3c66e71468f.png?imageMogr2/auto-orient/format/png/crop/!300x300a0a0
				Name      string `json:"Name"`      // coding
				Status    int64  `json:"Status"`    // 2
			} `json:"Watchers"`
			WorkingHours int64 `json:"WorkingHours"` // 0
		} `json:"Issue"`
		RequestID string `json:"RequestId"` // c2e8cdce-9d9f-be99-e209-3dae07b690ac
	} `json:"Response"`
}

// CreateIssue 创建事项
func (i *IssueClient) CreateIssue(req CreateIssueReq) (resp CreateIssueResp, err error) {
	err = i.Call(&req, &resp)
	return
}
