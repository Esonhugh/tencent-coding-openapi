package test

type DeleteReportReq struct {
	Action      string `json:"Action"`      // DeleteReport
	ProjectName string `json:"ProjectName"` // project name
	ReportID    int64  `json:"ReportId"`    // 1
}

func (req *DeleteReportReq) SetAction() string {
	req.Action = "DeleteReport"
	return req.Action
}

type DeleteReportResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 980cbd88-caac-626a-8c8d-8019acf41bb4
	} `json:"Response"`
}

// DeleteReport 删除测试报告
func (t *TestClient) DeleteReport(req DeleteReportReq) (resp DeleteReportResp, err error) {
	err = t.Call(&req, &resp)
	return
}
