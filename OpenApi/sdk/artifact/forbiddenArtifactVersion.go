package artifact

type ForbiddenArtifactVersionReq struct {
	Action          string `json:"Action"`          // ForbiddenArtifactVersion
	ForbiddenAction string `json:"ForbiddenAction"` // FORBIDDEN
	ForbiddenNote   string `json:"ForbiddenNote"`   // 垃圾制品，禁止下载
	Package         string `json:"Package"`         // Pro
	PackageVersion  string `json:"PackageVersion"`  // latest
	ProjectID       int64  `json:"ProjectId"`       // 450
	Repository      string `json:"Repository"`      // generic
}

func (req *ForbiddenArtifactVersionReq) SetAction() string {
	req.Action = "ForbiddenArtifactVersion"
	return req.Action
}

type ForbiddenArtifactVersionResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ForbiddenArtifactVersion 禁止、解禁制品版本下载
func (a *ArtifactClient) ForbiddenArtifactVersion(req ForbiddenArtifactVersionReq) (resp ForbiddenArtifactVersionResp, err error) {
	err = a.Call(&req, &resp)
	return
}
