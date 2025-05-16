package main

import (
	"mailx/cmd"
	"mailx/conf"
)

func main() {
	conf.InitConfig()
	conf.InitLogger()
	cmd.ServerRun()
}
