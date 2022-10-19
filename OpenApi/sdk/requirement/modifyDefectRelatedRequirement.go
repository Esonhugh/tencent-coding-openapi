package requirement

type ModifyDefectRelatedRequirementReq struct {
	DefectCode      int64  `json:"DefectCode"`      // 2
	ProjectName     string `json:"ProjectName"`     // project-demo
	RequirementCode int64  `json:"RequirementCode"` // 3
}

func (req *ModifyDefectRelatedRequirementReq) SetAction() string {
	return "ModifyDefectRelatedRequirement"

}

type ModifyDefectRelatedRequirementResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ModifyDefectRelatedRequirement 修改缺陷所属的需求
func (r *RequirementClient) ModifyDefectRelatedRequirement(req ModifyDefectRelatedRequirementReq) (resp ModifyDefectRelatedRequirementResp, err error) {
	err = r.Call(&req, &resp)
	return
}
