package add

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/cmd"
	"github.com/esonhugh/tencent-coding-openapi/service/config"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(SubCmd)
	SubCmd.AddCommand(SubCmdSsh)
}

var Client *OpenApi.Client

var SubCmd = &cobra.Command{
	Use:     "add",
	Short:   "添加对象 (Add Objects)",
	Long:    "添加对象 (Add Objects)",
	Example: "add <object>",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		token := config.GlobalConfig.GetString("auth.access_token")
		isOAuth := config.GlobalConfig.GetBool("auth.is_oauth")
		Client = OpenApi.NewClient()
		Client.SetToken(isOAuth, token)
	},
}
