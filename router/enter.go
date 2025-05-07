package router

import (
	"blogW_server/global"
	"blogW_server/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(global.Config.System.GinMode) // gin模式配置
	r := gin.Default()                        // 创建根路由

	r.Static("/uploads", "uploads") // 配置静态路由(映射)

	nr := r.Group("/api") // 创建子路由组 `/api`

	nr.Use(middleware.LogMiddleware) // 使用中间件
	SiteRouter(nr)                   // 在子路由组 `/api` 下注册 `/site` 路由
	LogRouter(nr)
	ImageRouter(nr)

	addr := global.Config.System.Addr()
	r.Run(addr)
}
