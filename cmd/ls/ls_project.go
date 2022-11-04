package ls

import (
	"fmt"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/team"
	"github.com/esonhugh/tencent-coding-openapi/config"
	"github.com/esonhugh/tencent-coding-openapi/utils/Error"
	"github.com/esonhugh/tencent-coding-openapi/utils/Print"
	"github.com/spf13/cobra"
)

var SubCmdProject = &cobra.Command{
	Use:   "project",
	Short: "列出项目 (List Projects)",
	Long:  "列出项目 (List Projects)",
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GlobalConfig.GetString("auth.access_token")
		isOAuth := config.GlobalConfig.GetBool("auth.is_oauth")
		Client := OpenApi.NewClient()
		Client.SetToken(isOAuth, token)
		resp, err := Client.DescribeCodingProjects(team.DescribeCodingProjectsReq{
			PageNumber: 1,
			PageSize:   20,
		})
		Error.HandleError(err)
		if resp.Response.Data.TotalCount > 20 {
			resp, err = Client.DescribeCodingProjects(team.DescribeCodingProjectsReq{
				PageNumber: 1,
				PageSize:   resp.Response.Data.TotalCount,
			})
		}
		Error.HandleError(err)

		var t Print.Table
		t.Header = []string{"ID", "Name", "DisplayName", "Description"}
		for _, each := range resp.Response.Data.ProjectList {
			t.Body = append(t.Body, []string{fmt.Sprintf("%v", each.ID), each.Name, each.DisplayName, each.Description})
		}
		t.Print("Project List")
	},
}
