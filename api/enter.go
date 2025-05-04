package api

import (
	"blogW_server/api/log_api"
	"blogW_server/api/site_api"
)

type Api struct {
	SiteApi site_api.SiteApi
	LogApi  log_api.LogApi
}

var App = Api{} // 自动初始化子模块结构体
