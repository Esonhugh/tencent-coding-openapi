package artifact

type ReleaseArtifactVersionReq struct {
	Action         string `json:"Action"`         // ReleaseArtifactVersion
	Package        string `json:"Package"`        // Pro
	PackageVersion string `json:"PackageVersion"` // latest
	ProjectID      int64  `json:"ProjectId"`      // 450
	Repository     string `json:"Repository"`     // generic
}

func (req *ReleaseArtifactVersionReq) SetAction() string {
	req.Action = "ReleaseArtifactVersion"
	return req.Action
}

type ReleaseArtifactVersionResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ReleaseArtifactVersion 发布制品版本
func (a *ArtifactClient) ReleaseArtifactVersion(req ReleaseArtifactVersionReq) (resp ReleaseArtifactVersionResp, err error) {
	err = a.Call(&req, &resp)
	return
}
