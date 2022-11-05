package ls

import (
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/service/config"
	"github.com/spf13/cobra"
)

var SubCmdRepo = &cobra.Command{
	Use:   "repo",
	Short: "列出仓库 (List Repositories)",
	Long:  "列出仓库 (List Repositories)",
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GlobalConfig.GetString("auth.access_token")
		isOAuth := config.GlobalConfig.GetBool("auth.is_oauth")
		Client := OpenApi.NewClient()
		Client.SetToken(isOAuth, token)

	},
}
