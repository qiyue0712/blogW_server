package router

import (
	"blogW_server/api"
	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.SiteApi          // 获取处理函数总api
	r.GET("site", app.SiteInfoView) // 路由路径 + 处理函数本身
	r.PUT("site", app.SiteUpdateView)
}
