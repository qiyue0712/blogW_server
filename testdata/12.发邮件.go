package main

import (
	"blogW_server/core"
	"blogW_server/flags"
	"blogW_server/global"
	"blogW_server/service/email_service"
)

func main() {
	flags.Parse()
	global.Config = core.Readconf()
	core.InitLogrus()

	email_service.SendRegisterCode("13641913653@163.com", "0712")
}
