package artifact

type CreateArtifactPropertiesReq struct {
	Action         string `json:"Action"`         // CreateArtifactProperties
	Package        string `json:"Package"`        // Pro
	PackageVersion string `json:"PackageVersion"` // latest
	ProjectID      int64  `json:"ProjectId"`      // 450
	PropertySet    []struct {
		Name  string `json:"Name"`  // name1
		Value string `json:"Value"` // value1
	} `json:"PropertySet"`
	Repository string `json:"Repository"` // generic
}

func (req *CreateArtifactPropertiesReq) SetAction() string {
	req.Action = "CreateArtifactProperties"
	return req.Action
}

type CreateArtifactPropertiesResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// CreateArtifactProperties 新增制品属性
func (a *ArtifactClient) CreateArtifactProperties(req CreateArtifactPropertiesReq) (resp CreateArtifactPropertiesResp, err error) {
	err = a.Call(&req, &resp)
	return
}
