package cmd

import (
	"github.com/esonhugh/tencent-coding-openapi/utils/log"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "codingcli",
	Short: "codingCli",
	Long: `Hi, There is CodingCli.
[+] Coding Cli is an Cli Tool which involved SDK implementation of Coding Platform APIs.
`,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		log.Init(logLevel)
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var logLevel, configFile string

func init() {
	RootCmd.PersistentFlags().StringVar(&logLevel, "logLevel", "info", "设置日志等级 (Set log level) [trace|debug|info|warn|error|fatal|panic]")
	RootCmd.PersistentFlags().StringVar(&configFile, "config", "coding.yaml", "设置全局配置文件 (Set config file)")
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  RootCmd,
		Headings: cc.HiGreen + cc.Underline,
		Commands: cc.Cyan + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Cyan + cc.Bold,
	})
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
