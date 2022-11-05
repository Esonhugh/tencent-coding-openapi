package ls

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/define"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/git"
	"github.com/esonhugh/tencent-coding-openapi/service/config"
	"github.com/esonhugh/tencent-coding-openapi/service/db"
	"github.com/esonhugh/tencent-coding-openapi/utils/Error"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var SubCmdRepo = &cobra.Command{
	Use:     "repo",
	Short:   "列出仓库 (List Repositories)",
	Long:    "列出仓库 (List Repositories)",
	Example: "coding-cli ls repo 项目名(Like正则)",
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GlobalConfig.GetString("auth.access_token")
		isOAuth := config.GlobalConfig.GetBool("auth.is_oauth")
		Client := OpenApi.NewClient()
		Client.SetToken(isOAuth, token)
		DB := db.GlobalDatabase.MainDB
		var TeamIds []int
		err := DB.Model(define.ProjectObject{}).Distinct("team_id").Find(&TeamIds).Error
		Error.HandleError(err)
		for _, TeamId := range TeamIds {
			resp, err := Client.DescribeTeamDepotInfoList(git.DescribeTeamDepotInfoListReq{
				TeamID: int64(TeamId),
			})
			Error.HandleError(err)
			p := define.ConvertDepotObjectList(resp.Response.DepotData.Depots)
			p.PrintSelf()
			log.Debug("Saving in Database")
			DB.Save(p)
			log.Debug("Saved complete")
		}
	},
}
