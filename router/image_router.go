package router

import (
	"blogW_server/api"
	"blogW_server/middleware"
	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	app := api.App.ImageApi // 获取处理函数总api
	r.POST("images", middleware.AuthMiddleware, app.ImageUploadView)
	r.POST("images/qiniu", middleware.AuthMiddleware, app.QiNiuGenToken)
	r.GET("images", middleware.AdminMiddleware, app.ImageListView)
	r.DELETE("images", middleware.AdminMiddleware, app.ImageRemoveView)
}
