package config

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/esonhugh/tencent-coding-openapi/cmd"
	"github.com/esonhugh/tencent-coding-openapi/config"
	"github.com/esonhugh/tencent-coding-openapi/utils/Error"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(SubCommand)
}

const (
	PersonalToken = "个人 Token (PersonalToken)"
	OAuthToken    = "OAuth Token (OAuthToken)"
)

var SubCommand = &cobra.Command{
	Use:   "config",
	Short: "配置 AccessToken (config AccessToken)",
	Long:  "配置 AccessToken (config AccessToken)",
	Run: func(cmd *cobra.Command, args []string) {
		var TokenType string
		err := survey.AskOne(&survey.Select{
			Message: "是否是 OAuth Token (Is OAuth AccessToken)",
			Options: []string{PersonalToken, OAuthToken},
		}, &TokenType)
		Error.HandleFatal(err)

		config.GlobalConfig.Set("auth.IsOAuth", TokenType == OAuthToken)
		var Token string
		err = survey.AskOne(&survey.Input{
			Message: "输入 Token (Token)",
		}, &Token)
		Error.HandleFatal(err)

		config.GlobalConfig.Set("auth.AccessToken", Token)
		err = config.Save()
		Error.HandleFatal(err)
		log.Println("保存成功 (Save Successfully)")
	},
}
