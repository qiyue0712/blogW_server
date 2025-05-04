package site_api

import (
	"blogW_server/common/res"
	"github.com/gin-gonic/gin"
)

// 定义站点相关的 API 结构体

type SiteApi struct {
}

// 示例处理函数：获取站点信息

func (SiteApi) SiteInfoView(c *gin.Context) {
	res.OkWithData("xx", c)
	return
}

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
