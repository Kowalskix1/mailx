package main

import (
	"mailtrap/cmd"
	"mailtrap/conf"
)

func main() {
	conf.InitConfig()
	conf.InitLogger()
	cmd.ServerRun()
}
