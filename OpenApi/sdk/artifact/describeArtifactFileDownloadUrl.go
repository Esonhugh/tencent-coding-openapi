package artifact

type DescribeArtifactFileDownloadUrlReq struct {
	Action         string `json:"Action"`         // DescribeArtifactFileDownloadUrl
	FileName       string `json:"FileName"`       // Pro
	Package        string `json:"Package"`        // Pro
	PackageVersion string `json:"PackageVersion"` // latest
	ProjectID      int64  `json:"ProjectId"`      // 450
	Repository     string `json:"Repository"`     // generic
	Timeout        int64  `json:"Timeout"`        // 300
}

func (req *DescribeArtifactFileDownloadUrlReq) SetAction() string {
	req.Action = "DescribeArtifactFileDownloadUrl"
	return req.Action
}

type DescribeArtifactFileDownloadUrlResp struct {
	Response struct {
		RequestID int64  `json:"RequestId,string"` // 1
		URL       string `json:"Url"`              // https//xxxxxxx.com/xxx
	} `json:"Response"`
}

// DescribeArtifactFileDownloadUrl 查询制品文件临时下载链接
func (a *ArtifactClient) DescribeArtifactFileDownloadUrl(req DescribeArtifactFileDownloadUrlReq) (resp DescribeArtifactFileDownloadUrlResp, err error) {
	err = a.Call(&req, &resp)
	return
}
