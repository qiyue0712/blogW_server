package site_api

import (
	"blogW_server/common/res"
	"blogW_server/global"
	"blogW_server/middleware"
	"github.com/gin-gonic/gin"
)

// 定义站点相关的 API 结构体

type SiteApi struct {
}

type SiteInfoRequest struct {
	Name string `uri:"name"`
}

// 示例处理函数：获取站点信息

func (SiteApi) SiteInfoView(c *gin.Context) {
	var cr SiteInfoRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	if cr.Name == "site" {
		global.Config.Site.About.Version = global.Version
		res.OkWithData(global.Config.Site, c)
		return
	}

	// 判断角色是不是管理员
	middleware.AdminMiddleware(c)
	_, ok := c.Get("claims")
	if !ok {
		return
	}

	var data any

	switch cr.Name {
	case "email":
		rep := global.Config.Email
		rep.AuthCode = "******"
		data = rep
	case "qq":
		rep := global.Config.QQ
		rep.AppKey = "******"
		data = rep
	case "qiNiu":
		rep := global.Config.QiNiu
		rep.SecretKey = "******"
		data = rep
	case "ai":
		rep := global.Config.Ai
		rep.SecretKey = "******"
		data = rep
	default:
		res.FailWithMsg("不存在的配置", c)
		return
	} // 站点配置信息接口

	res.OkWithData(data, c)
	return
}

func (SiteApi) SiteInfoQQView(c *gin.Context) {
	res.OkWithData(global.Config.QQ.Url(), c)
} // qq登录接口

// 解析

type SiteUpDateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required" label:"年龄"`
}

// 更新站点

func (SiteApi) SiteUpdateView(c *gin.Context) {
	//log := log_service.GetLog(c) // 由中间件传递的log

	// 请求体的获取
	var cr SiteUpDateRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	res.OkWithData("更新成功", c)
	return
}
