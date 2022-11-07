package ls

import (
	"github.com/esonhugh/tencent-coding-openapi/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(SubCmd)
	SubCmd.AddCommand(SubCmdProject)
	SubCmd.AddCommand(SubCmdRepo)
	SubCmd.AddCommand(SubCmdMember)
}

// SubCmd is core cobra.Command of subcommand
var SubCmd = &cobra.Command{
	Use:   "ls",
	Short: "列出对象 (List Objects)",
	Long: "列出对象 (List Objects)" +
		"ls 中获得的内容将会缓存到 cache.sqlite (默认的 cache 数据库) 文件中" +
		"如果有需要可以直接使用数据库链接软件打开数据库",
	Example: "ls <object>",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}
