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

var SubCmdMember = &cobra.Command{
	Use:   "member",
	Short: "列出成员 (List Members)",
	Long:  "列出成员 (List Members)",
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GlobalConfig.GetString("auth.access_token")
		isOAuth := config.GlobalConfig.GetBool("auth.is_oauth")
		Client := OpenApi.NewClient()
		Client.SetToken(isOAuth, token)
		resp, err := Client.DescribeTeamMembers(team.DescribeTeamMembersReq{
			PageNumber: 1,
			PageSize:   10,
		})
		Error.HandleError(err)
		if resp.Response.Data.TotalCount > 10 {
			resp, err = Client.DescribeTeamMembers(team.DescribeTeamMembersReq{
				PageNumber: 1,
				PageSize:   resp.Response.Data.TotalCount,
			})
		}
		Error.HandleError(err)
		p := define.ConvertMemberObjectList(resp.Response.Data.TeamMembers)
		p.PrintSelf()

		log.Debug("Saving in Database")
		DB := db.GlobalDatabase.MainDB.Save(p)
		Error.HandleError(DB.Error)
		log.Debug("Saved complete")
	},
}
