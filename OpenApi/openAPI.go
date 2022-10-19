package OpenApi

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/baseClient"
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/artifact"         // 制品仓库
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/branch"           // 分支
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/branchProtection" // 保护分支规则
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/cd"               // 持续部署
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/ci"               // 持续集成
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/commit"           // 提交
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/git"              // 代码托管
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/issue"            // 事项
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/iteration"        // 迭代
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/mergeReq"         // 合并请求
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/project"          // 项目
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/projectSetting"   // 项目设置
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/protectedBranch"  // 保护分支
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/release"          // 版本
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/requirement"      // 需求
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/tag"              // 标签
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/team"             // 团队
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/test"             // 测试管理
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/wiki"             // 文档管理
	"github.com/esonhugh/tencent-coding-openapi/openAPI/sdk/work"             // 工时
)

type Client struct {
	*baseClient.ApiClientBase
	team.TeamClient                         // 团队
	project.ProjectClient                   // 项目
	projectSetting.ProjectSettingClient     // 项目设置
	requirement.RequirementClient           // 需求
	issue.IssueClient                       // 事项
	iteration.IterationClient               // 迭代
	work.WorkClient                         // 工时
	git.GitClient                           // 代码托管
	commit.CommitClient                     // 提交
	tag.TagClient                           // 标签
	release.ReleaseClient                   // 版本
	branch.BranchClient                     // 分支
	protectedBranch.ProtectedBranchClient   // 保护分支
	branchProtection.BranchProtectionClient // 保护分支规则
	mergeReq.MergeReqClient                 // 合并请求
	ci.CiClient                             // 持续集成
	cd.CdClient                             // 持续部署
	artifact.ArtifactClient                 // 制品仓库
	test.TestClient                         // 测试管理
	wiki.WikiClient                         // 文档管理
}

func NewClient() *Client {
	base := baseClient.CreateApiClient(nil)
	return &Client{
		base,
		team.New(base),
		project.New(base),
		projectSetting.New(base),
		requirement.New(base),
		issue.New(base),
		iteration.New(base),
		work.New(base),
		git.New(base),
		commit.New(base),
		tag.New(base),
		release.New(base),
		branch.New(base),
		protectedBranch.New(base),
		branchProtection.New(base),
		mergeReq.New(base),
		ci.New(base),
		cd.New(base),
		artifact.New(base),
		test.New(base),
		wiki.New(base),
	}
}

func (c *Client) SetToken(isOAuth bool, token string) {
	c.SetAccessToken(isOAuth, token)
}
