package version

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/define"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/projectSetting"
	"github.com/esonhugh/tencent-coding-openapi/service/config"
	"github.com/esonhugh/tencent-coding-openapi/service/db"
	"github.com/esonhugh/tencent-coding-openapi/utils/Error"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var SubCmdMyProject = &cobra.Command{
	Use:   "project",
	Short: "列出自己所在项目 (List self Projects)",
	Long:  "列出自己所在项目 (List self Projects)",
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GlobalConfig.GetString("auth.access_token")
		isOAuth := config.GlobalConfig.GetBool("auth.is_oauth")
		Client := OpenApi.NewClient()
		Client.SetToken(isOAuth, token)
		resp, err := Client.GetMe()
		Error.HandleError(err)
		respProjects, err := Client.DescribeUserProjects(projectSetting.DescribeUserProjectsReq{
			UserID: int64(resp.ID),
		})
		Error.HandleError(err)
		p := define.ConvertProjectObjectList(respProjects.Response.ProjectList)
		p.PrintSelf()
		log.Debug("Saving in Database")
		DB := db.GlobalDatabase.MainDB.Save(p)
		Error.HandleError(DB.Error)
		log.Debug("Saved complete")
	},
}
