package main

import (
	"blogW_server/core"
	"blogW_server/flags"
	"blogW_server/global"
	"blogW_server/service/log_service"
)

func main() {
	flags.Parse()                   // 环境变量参数
	global.Config = core.Readconf() // 读配置文件
	core.InitLogrus()               // 日志初始化
	global.DB = core.InitDB()

	log := log_service.NewRuntimeLog("同步文章数据", log_service.RuntimeDateHour)
	log.SetItem("文章1", 11)
	log.Save()
	log.SetItem("文章2", 12)
	log.Save()
}
