package version

import (
	"errors"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/define"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/projectSetting"
	"github.com/esonhugh/tencent-coding-openapi/service/db"
	"github.com/esonhugh/tencent-coding-openapi/utils/Error"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var SubCmdMyProject = &cobra.Command{
	Use:   "project",
	Short: "列出自己所在项目 (List self Projects)",
	Long:  "列出自己所在项目 (List self Projects)",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := Client.GetMe()
		Error.HandleError(err)
		respProjects, err := Client.DescribeUserProjects(projectSetting.DescribeUserProjectsReq{
			UserID: int64(resp.ID),
		})
		Error.HandleError(err)
		if len(respProjects.Response.ProjectList) <= 0 {
			errors.New("count is less then one")
			os.Exit(1)
		}
		p := define.ConvertProjectObjectList(respProjects.Response.ProjectList)
		p.PrintSelf()
		log.Debug("Saving in Database")
		DB := db.GlobalDatabase.MainDB.Save(p)
		Error.HandleError(DB.Error)
		log.Debug("Saved complete")
	},
}
