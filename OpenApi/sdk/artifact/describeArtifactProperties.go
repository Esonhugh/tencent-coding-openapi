package artifact

type DescribeArtifactPropertiesReq struct {
	Action         string `json:"Action"`         // DescribeArtifactProperties
	Package        string `json:"Package"`        // Pro
	PackageVersion string `json:"PackageVersion"` // latest
	PageNumber     int64  `json:"PageNumber"`     // 1
	PageSize       int64  `json:"PageSize"`       // 10
	ProjectID      int64  `json:"ProjectId"`      // 450
	Repository     string `json:"Repository"`     // generic
}

func (req *DescribeArtifactPropertiesReq) SetAction() string {
	req.Action = "DescribeArtifactProperties"
	return req.Action
}

type DescribeArtifactPropertiesResp struct {
	Response struct {
		InstanceSet []struct {
			CreatedAt int64   `json:"CreatedAt"`      // 1.538743856753e+12
			ID        int64   `json:"Id"`             // 1
			Immutable bool    `json:"Immutable"`      // true
			Name      string  `json:"Name"`           // demo
			Value     string  `json:"Value"`          // demoValue
			Version   float64 `json:"Version,string"` // 2.0
		} `json:"InstanceSet"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeArtifactProperties 查询制品属性列表
func (a *ArtifactClient) DescribeArtifactProperties(req DescribeArtifactPropertiesReq) (resp DescribeArtifactPropertiesResp, err error) {
	err = a.Call(&req, &resp)
	return
}
