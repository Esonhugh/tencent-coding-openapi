package artifact

type DeleteArtifactPropertiesReq struct {
	Action          string   `json:"Action"`          // DeleteArtifactProperties
	Package         string   `json:"Package"`         // Pro
	PackageVersion  string   `json:"PackageVersion"`  // latest
	ProjectID       int64    `json:"ProjectId"`       // 450
	PropertyNameSet []string `json:"PropertyNameSet"` // name1
	Repository      string   `json:"Repository"`      // generic
}

func (req *DeleteArtifactPropertiesReq) SetAction() string {
	req.Action = "DeleteArtifactProperties"
	return req.Action
}

type DeleteArtifactPropertiesResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DeleteArtifactProperties 删除制品属性
func (a *ArtifactClient) DeleteArtifactProperties(req DeleteArtifactPropertiesReq) (resp DeleteArtifactPropertiesResp, err error) {
	err = a.Call(&req, &resp)
	return
}
