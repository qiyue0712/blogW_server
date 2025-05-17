package main

import (
	"blogW_server/core"
	"blogW_server/flags"
	"blogW_server/global"
	"blogW_server/service/qq_service"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.Readconf()
	core.InitLogrus()
	fmt.Println(qq_service.GetUserInfo(""))
}
