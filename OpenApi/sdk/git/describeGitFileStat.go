package git

type DescribeGitFileStatReq struct {
	Action    string `json:"Action"`    // DescribeGitFileStat
	DepotPath string `json:"DepotPath"` // codingcorp/k8s/etcd
	Path      string `json:"Path"`      // foo.txt
	Ref       string `json:"Ref"`       // master
}

func (req *DescribeGitFileStatReq) SetAction() string {
	req.Action = "DescribeGitFileStat"
	return req.Action
}

type DescribeGitFileStatResp struct {
	Response struct {
		Payload struct {
			IsExist bool `json:"IsExist"` // true
		} `json:"Payload"`
		RequestID string `json:"RequestId"` // 0d5ffa38-6180-4eb6-b957-b041dde33ba9
	} `json:"Response"`
}

// DescribeGitFileStat 查询某分支下某文件是否存在
func (g *GitClient) DescribeGitFileStat(req DescribeGitFileStatReq) (resp DescribeGitFileStatResp, err error) {
	err = g.Call(&req, &resp)
	return
}
