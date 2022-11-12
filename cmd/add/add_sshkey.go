package add

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi"
	"github.com/esonhugh/tencent-coding-openapi/OpenApi/sdk/git"
	"github.com/esonhugh/tencent-coding-openapi/service/config"
	"github.com/esonhugh/tencent-coding-openapi/utils/Error"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var SubCmdSsh = &cobra.Command{
	Use:   "ssh",
	Short: "添加 SSH Key (Add SSH Key)",
	Long:  "添加 SSH Key (Add SSH Key)",
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GlobalConfig.GetString("auth.access_token")
		isOAuth := config.GlobalConfig.GetBool("auth.is_oauth")
		name := config.GlobalConfig.GetString("ssh.name")
		sshKey := config.GlobalConfig.GetString("ssh.public_key")
		for sshKey == "" {
			log.Info("请先配置 SSH Key")
			err := survey.AskOne(&survey.Input{
				Message: "请输入 SSH Public Key (Please input your SSH Key):",
			}, &sshKey)
			Error.HandleFatal(err)
			if sshKey != "" {
				err = survey.AskOne(&survey.Input{
					Message: "请输入 SSH Key 名称 (Please input your SSH Key name):",
				}, &name)
				Error.HandleError(err)
				config.GlobalConfig.Set("ssh.public_key", sshKey)
				config.GlobalConfig.Set("ssh.name", name)
				Error.HandleError(config.Save())
			}
		}
		Client := OpenApi.NewClient()
		Client.SetToken(isOAuth, token)
		resp, err := Client.CreateSshKey(git.CreateSshKeyReq{
			Title:   name,
			Content: sshKey,
		})
		Error.HandlePanic(err)
		log.Info("添加成功 (Add Successfully)")
		log.Info("requestID: ", resp.Response.RequestID)
		log.Info("key finger print: ", resp.Response.SshKeyInfo.FingerPrint)
	},
}
