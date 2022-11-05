package ls

import (
	"github.com/esonhugh/tencent-coding-openapi/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(SubCmd)
	SubCmd.AddCommand(SubCmdProject)
	SubCmd.AddCommand(SubCmdRepo)
}

// SubCmd is core cobra.Command of subcommand
var SubCmd = &cobra.Command{
	Use:     "ls",
	Short:   "列出对象 (List Objects)",
	Long:    "列出对象 (List Objects)",
	Example: "ls <object>",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}
