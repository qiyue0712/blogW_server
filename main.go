package main

import (
	"blogW_server/core"
	"blogW_server/flags"
	"blogW_server/global"
)

func main() {
	flags.Parse()                   // 环境变量参数
	global.Config = core.Readconf() //读配置文件
	core.InitLogrus()               //日志初始化
	global.DB = core.InitDB()       //连接数据库
}
