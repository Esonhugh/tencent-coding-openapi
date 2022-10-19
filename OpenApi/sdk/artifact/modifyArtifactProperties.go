package artifact

type ModifyArtifactPropertiesReq struct {
	Action         string `json:"Action"`         // ModifyArtifactProperties
	Package        string `json:"Package"`        // Pro
	PackageVersion string `json:"PackageVersion"` // latest
	ProjectID      int64  `json:"ProjectId"`      // 450
	PropertySet    []struct {
		Name  string `json:"Name"`  // name1
		Value string `json:"Value"` // newValue
	} `json:"PropertySet"`
	Repository string `json:"Repository"` // generic
}

func (req *ModifyArtifactPropertiesReq) SetAction() string {
	req.Action = "ModifyArtifactProperties"
	return req.Action
}

type ModifyArtifactPropertiesResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ModifyArtifactProperties 修改制品属性
func (a *ArtifactClient) ModifyArtifactProperties(req ModifyArtifactPropertiesReq) (resp ModifyArtifactPropertiesResp, err error) {
	err = a.Call(&req, &resp)
	return
}
