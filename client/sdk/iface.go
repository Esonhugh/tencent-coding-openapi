package sdk

import (
	"github.com/esonhugh/tencent-coding-openapi/client/sdk/CommonSDK"
	"github.com/esonhugh/tencent-coding-openapi/client/sdk/TeamSDK"
)

type SDK interface {
	CommonSDK.CommonSDK
	TeamSDK.TeamSDK
	ProjectSDK
}

type ProjectSDK interface {
	CreateCodingProject()
	ModifyProject()
	DeleteOneProject()
	DescribeOneProject()
	DescribeProjectMembers()
	DescribeProjectByName()
	CreateProjectMember()
}

type ProjectSettingSDK interface {
	DescribeUserProjects()
	DescribeCodingCurrentUser()
	ModifyProjectPermission()
	DescribeProjectRoles()
}

type RequirementSDK interface {
	DescribeRequirementDefectRelation()
	ModifyDefectRelatedRequirement()
	CreateRequirementDefectRelation()
}

type IssueSDK interface {
	CreateIssue()
	DeleteIssue()
	// Todo: To Be Continued..
}

// Todo: Tobe Continued...
