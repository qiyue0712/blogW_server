package router

import (
	"blogW_server/api"
	"blogW_server/middleware"
	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.SiteApi // 获取处理函数总api
	r.GET("site/qq_url", app.SiteInfoQQView)
	r.GET("site/:name", app.SiteInfoView) // 路由路径 + 处理函数本身 站点信息

	r.PUT("site/:name", middleware.AdminMiddleware, app.SiteUpdateView) // 网站配置更新
}
