package test

type DescribeReportReq struct {
	Action      string `json:"Action"`      // DescribeReport
	ProjectName string `json:"ProjectName"` // project name
	ReportID    int64  `json:"ReportId"`    // 1
}

func (req *DescribeReportReq) SetAction() string {
	req.Action = "DescribeReport"
	return req.Action
}

type DescribeReportResp struct {
	Response struct {
		Data struct {
			Report struct {
				CreatedAt      string `json:"CreatedAt"`           // 2021-06-23 17:10:51
				CreatedBy      int64  `json:"CreatedBy"`           // 1
				ID             int64  `json:"Id"`                  // 1595
				IterationIDs   int64  `json:"IterationIds,string"` // 1
				IterationNames string `json:"IterationNames"`      // Sprint-1
				Name           string `json:"Name"`                // 报告
				ProjectName    string `json:"ProjectName"`         // Project Name
				ReportOverview struct {
					AutomationPercent       int64 `json:"AutomationPercent"`       // 0
					AvgClosedSeconds        int64 `json:"AvgClosedSeconds"`        // -1
					CaseSum                 int64 `json:"CaseSum"`                 // 3
					CompletedSum            int64 `json:"CompletedSum"`            // 0
					DefectFixPercent        int64 `json:"DefectFixPercent"`        // 0
					DefectReopenPercent     int64 `json:"DefectReopenPercent"`     // 0
					DefectSum               int64 `json:"DefectSum"`               // 0
					DurationFixed           int64 `json:"DurationFixed"`           // -1
					ExecPercent             int64 `json:"ExecPercent"`             // 100
					IssuesSum               int64 `json:"IssuesSum"`               // 4
					PassPercent             int64 `json:"PassPercent"`             // 0
					ProcessingSum           int64 `json:"ProcessingSum"`           // 0
					RequirementCoverPercent int64 `json:"RequirementCoverPercent"` // 50
					TodoSum                 int64 `json:"TodoSum"`                 // 4
				} `json:"ReportOverview"`
				RunIDs              []string    `json:"RunIds"`              // 20
				RunNames            []string    `json:"RunNames"`            // 计划
				StatisticsEndTime   string      `json:"StatisticsEndTime"`   // 2021-06-23
				StatisticsStartTime string      `json:"StatisticsStartTime"` // 2021-06-22
				Status              string      `json:"Status"`              // AVAILABLE
				Summary             interface{} `json:"Summary"`             // <nil>
				TemplateID          interface{} `json:"TemplateId"`          // <nil>
			} `json:"Report"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 980cbd88-caac-626a-8c8d-8019acf41bb4
	} `json:"Response"`
}

// DescribeReport 测试报告详情
func (t *TestClient) DescribeReport(req DescribeReportReq) (resp DescribeReportResp, err error) {
	err = t.Call(&req, &resp)
	return
}
