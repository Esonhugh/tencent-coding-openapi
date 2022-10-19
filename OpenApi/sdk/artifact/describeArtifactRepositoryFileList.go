package artifact

type DescribeArtifactRepositoryFileListReq struct {
	Action     string `json:"Action"`     // DescribeArtifactRepositoryFileList
	PageSize   int64  `json:"PageSize"`   // 1000
	Project    string `json:"Project"`    // testing
	Repository string `json:"Repository"` // maven
}

func (req *DescribeArtifactRepositoryFileListReq) SetAction() string {
	req.Action = "DescribeArtifactRepositoryFileList"
	return req.Action
}

type DescribeArtifactRepositoryFileListResp struct {
	Response struct {
		Data struct {
			ContinuationToken string `json:"ContinuationToken"` // 5e739c1e1f8bdb20cdb0db2ab2215903
			InstanceSet       []struct {
				ArtifactType int64  `json:"ArtifactType"` // 3
				DownloadURL  string `json:"DownloadUrl"`  // https://coding.net/repository/project/maven/pkg/vers/folder/object-1.jar
				Host         string `json:"Host"`         // coding.net
				Path         string `json:"Path"`         // folder/object-1.jpg
				Project      string `json:"Project"`      // project
				Repository   string `json:"Repository"`   // maven
			} `json:"InstanceSet"`
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeArtifactRepositoryFileList 查询制品仓库可下载文件列表
func (a *ArtifactClient) DescribeArtifactRepositoryFileList(req DescribeArtifactRepositoryFileListReq) (resp DescribeArtifactRepositoryFileListResp, err error) {
	err = a.Call(&req, &resp)
	return
}
