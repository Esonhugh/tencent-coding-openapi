package ls

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/define"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/team"
	"github.com/esonhugh/tencent-coding-openapi/service/config"
	"github.com/esonhugh/tencent-coding-openapi/service/db"
	"github.com/esonhugh/tencent-coding-openapi/utils/Error"
	log "github.com/sirupsen/logrus"
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
		p := define.ConvertProjectObjectList(resp.Response.Data.ProjectList)
		p.PrintSelf()
		log.Debug("Saving in Database")
		DB := db.GlobalDatabase.MainDB.Save(p)
		Error.HandleError(DB.Error)
		log.Debug("Saved complete")
	},
}
