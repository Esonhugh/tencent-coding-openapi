package artifact

type CreateArtifactRepositoryReq struct {
	AccessLevel    int64  `json:"AccessLevel"`    // 1
	Action         string `json:"Action"`         // CreateArtifactRepository
	AllowProxy     bool   `json:"AllowProxy"`     // false
	ProjectID      int64  `json:"ProjectId"`      // 450
	RepositoryName string `json:"RepositoryName"` // api_created
	Type           int64  `json:"Type"`           // 1
}

func (req *CreateArtifactRepositoryReq) SetAction() string {
	req.Action = "CreateArtifactRepository"
	return req.Action
}

type CreateArtifactRepositoryResp struct {
	Response struct {
		ID        int64 `json:"Id"`               // 1
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// CreateArtifactRepository 创建制品仓库
func (a *ArtifactClient) CreateArtifactRepository(req CreateArtifactRepositoryReq) (resp CreateArtifactRepositoryResp, err error) {
	err = a.Call(&req, &resp)
	return
}
