package add

import (
	"github.com/esonhugh/tencent-coding-openapi/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(SubCmd)
	SubCmd.AddCommand(SubCmdSsh)
}

var SubCmd = &cobra.Command{
	Use:     "add",
	Short:   "添加对象 (Add Objects)",
	Long:    "添加对象 (Add Objects)",
	Example: "add <object>",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}
