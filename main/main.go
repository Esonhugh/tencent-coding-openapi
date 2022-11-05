package main

import (
	"github.com/esonhugh/tencent-coding-openapi/cmd"
	_ "github.com/esonhugh/tencent-coding-openapi/cmd/config"
	_ "github.com/esonhugh/tencent-coding-openapi/cmd/ls"
	_ "github.com/esonhugh/tencent-coding-openapi/cmd/me"
	_ "github.com/esonhugh/tencent-coding-openapi/cmd/version" // import sub command as module
)

func init() {
}

func main() {
	cmd.Execute()
}
