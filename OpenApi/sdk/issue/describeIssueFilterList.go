package issue

type DescribeIssueFilterListReq struct {
	Action      string `json:"Action"`      // DescribeIssueFilterList
	IssueType   string `json:"IssueType"`   // REQUIREMENT
	ProjectName string `json:"ProjectName"` // demo-project
}

func (req *DescribeIssueFilterListReq) SetAction() string {
	req.Action = "DescribeIssueFilterList"
	return req.Action
}

type DescribeIssueFilterListResp struct {
	Response struct {
		Data struct {
			CustomFilterList []interface{} `json:"CustomFilterList"` // <nil>
			SystemFilterList []struct {
				Content    string `json:"Content"`    // {"FilterIssueType":"REQUIREMENT","Sort":{"Key":"PRIORITY","Value":"DESC"},"Conditions":[{"Key":"STATUS_TYPE","Fixed":true,"ConstValue":null,"FilterIssueType":"REQUIREMENT","ProjectId":1,"ProjectIds":null,"Value":["TODO","PROCESSING"],"TeamView":false},{"Key":"ASSIGNEE","Fixed":true,"ConstValue":null,"FilterIssueType":"REQUIREMENT","ProjectId":1,"ProjectIds":null,"Value":[1],"ValueChanged":null,"ValidInfo":null,"Status":null,"TeamView":false}]}
				CreatedAt  int64  `json:"CreatedAt"`  // 1.597908832e+12
				CreatorID  int64  `json:"CreatorId"`  // 1
				ID         int64  `json:"Id"`         // 1
				IsDefault  bool   `json:"IsDefault"`  // false
				IsSystem   bool   `json:"IsSystem"`   // true
				IssueType  string `json:"IssueType"`  // REQUIREMENT
				Name       string `json:"Name"`       // 我未完成的
				SharedTeam bool   `json:"SharedTeam"` // false
				UpdatedAt  int64  `json:"UpdatedAt"`  // 1.597908832e+12
			} `json:"SystemFilterList"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 146e1c15-b712-a7b6-851b-16267a079fad
	} `json:"Response"`
}

// DescribeIssueFilterList 查询筛选器列表
func (i *IssueClient) DescribeIssueFilterList(req DescribeIssueFilterListReq) (resp DescribeIssueFilterListResp, err error) {
	err = i.Call(&req, &resp)
	return
}
