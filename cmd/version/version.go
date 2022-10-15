package version

import (
	"github.com/esonhugh/tencent-coding-openapi/cmd"
	"github.com/esonhugh/tencent-coding-openapi/utils/Print"
	"github.com/spf13/cobra"
)

var version = "0.1"

// version module

// init in modules will add self to RootCmd when init package.
func init() {
	cmd.RootCmd.AddCommand(SubCmd)
}

// SubCmd is core cobra.Command of subcommand
var SubCmd = &cobra.Command{
	Use:   "version",
	Short: "输出版本 (Print the version number)",
	Long:  "输出版本 (Print the version number)",
	Run: func(cmd *cobra.Command, args []string) {
		data := [][]string{
			{version, "Esonhugh"},
		}
		var header = []string{"当前版本 (Version)", "作者 (Author)"}
		var td = Print.Table{
			Header: header,
			Body:   data}
		td.Print("")
	},
}
