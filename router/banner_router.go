package router

import (
	"blogW_server/api"
	"blogW_server/middleware"
	"github.com/gin-gonic/gin"
)

func BannerRouter(r *gin.RouterGroup) {
	app := api.App.BannerApi // 获取处理函数总api
	r.GET("banner", app.BannerListView)
	r.POST("banner", middleware.AdminMiddleware, app.BannerCreateView)
	r.PUT("banner/:id", middleware.AdminMiddleware, app.BannerUpdateView)
	r.DELETE("banner", middleware.AdminMiddleware, app.BannerRemoveView)
}
