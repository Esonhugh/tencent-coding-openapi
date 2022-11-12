package version

import (
	"fmt"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/cmd"
	"github.com/esonhugh/tencent-coding-openapi/service/config"
	"github.com/esonhugh/tencent-coding-openapi/utils/Error"
	"github.com/esonhugh/tencent-coding-openapi/utils/Print"
	"github.com/spf13/cobra"
)

// init in modules will add self to RootCmd when init package.
func init() {
	cmd.RootCmd.AddCommand(SubCmd)
	SubCmd.AddCommand(SubCmdMyProject)
}

var Client *OpenApi.Client

// SubCmd is core cobra.Command of subcommand
var SubCmd = &cobra.Command{
	Use:   "me",
	Short: "个人信息 (Print the Personal Info)",
	Long:  "个人信息 (Print the Personal Info)",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		token := config.GlobalConfig.GetString("auth.access_token")
		isOAuth := config.GlobalConfig.GetBool("auth.is_oauth")
		Client = OpenApi.NewClient()
		Client.SetToken(isOAuth, token)
	},
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := Client.GetMe()
		Error.HandleError(err)
		var t Print.Table
		t.Header = []string{"姓名 (Name)", "邮箱 (Email)", "用户名 (Username)", "创建时间 (Created At)", "团队 (Team)"}
		t.Body = append(t.Body, []string{
			resp.Name, resp.Email, fmt.Sprintf("%v", resp.ID), fmt.Sprintf("%v", resp.CreateAt), resp.Team,
		})
		t.Print("")
	},
}
