package requirement

type CreateRequirementDefectRelationReq struct {
	DefectCode      int64  `json:"DefectCode"`      // 2
	ProjectName     string `json:"ProjectName"`     // project-demo
	RequirementCode int64  `json:"RequirementCode"` // 1
}

func (req *CreateRequirementDefectRelationReq) SetAction() string {
	return "CreateRequirementDefectRelation"
}

type CreateRequirementDefectRelationResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// CreateRequirementDefectRelation 需求关联缺陷
func (r *RequirementClient) CreateRequirementDefectRelation(req CreateRequirementDefectRelationReq) (resp CreateRequirementDefectRelationResp, err error) {
	err = r.Call(&req, &resp)
	return
}
