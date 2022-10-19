package issue

type DescribeIssueFileUrlReq struct {
	Action      string `json:"Action"`      // DescribeIssueFileUrl
	FileID      int64  `json:"FileId"`      // 7
	ProjectName string `json:"ProjectName"` // project-demo
}

func (req *DescribeIssueFileUrlReq) SetAction() string {
	req.Action = "DescribeIssueFileUrl"
	return req.Action
}

type DescribeIssueFileUrlResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 3cdcc77c-f1a3-1695-fafe-6bb4fc71b647
		URL       string `json:"Url"`       // http://coding-net-test-self-1257242599.cos.ap-shanghai.myqcloud.com/b6357cc0-e357-11ea-af4e-4177fbfe7a25.txt?sign=q-sign-algorithm=sha1&q-ak=AKIDUnFWaPTLQccQyIsuIdCdq04B8xd6iOK4&q-sign-time=1599013747;1599017347&q-key-time=1599013747;1599017347&q-header-list=&q-url-param-list=response-content-disposition;response-expires&q-signature=4aa5d8f991c0388561f0c43e0f15059e91ccede7&response-content-disposition=attachment;filename=testfile-20200821104044.txt&response-expires=Fri, 04 Sep 2020 14:29:07 GMT
	} `json:"Response"`
}

// DescribeIssueFileUrl 查询事项附件的下载地址
func (i *IssueClient) DescribeIssueFileUrl(req DescribeIssueFileUrlReq) (resp DescribeIssueFileUrlResp, err error) {
	err = i.Call(&req, &resp)
	return
}
