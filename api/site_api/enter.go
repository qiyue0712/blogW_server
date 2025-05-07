package site_api

import (
	"blogW_server/common/res"
	"blogW_server/conf"
	"blogW_server/core"
	"blogW_server/global"
	"blogW_server/middleware"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
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
	var cr SiteInfoRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	// 网站配置更新
	var rep any
	switch cr.Name {
	case "site":
		var data conf.Site
		err = c.ShouldBindJSON(&data)
		rep = data
	case "email":
		var data conf.Email
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qq":
		var data conf.QQ
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qiu":
		var data conf.QiNiu
		err = c.ShouldBindJSON(&data)
		rep = data
	case "ai":
		var data conf.Ai
		err = c.ShouldBindJSON(&data)
		rep = data
	default:
		res.FailWithMsg("不存在的配置", c)
		return
	}
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	//断言数据类型
	switch s := rep.(type) {
	case conf.Site:
		// 判断站点信息更新前端文件部分
		err = UpdateSite(s)
		if err != nil {
			res.FailWithError(err, c)
			return
		}

		global.Config.Site = s
	case conf.Email:
		if s.AuthCode == "******" {
			s.AuthCode = global.Config.Email.AuthCode
		}
		global.Config.Email = s
	case conf.QQ:
		if s.AppKey == "******" {
			s.AppKey = global.Config.QQ.AppKey
		}
		global.Config.QQ = s
	case conf.QiNiu:
		if s.SecretKey == "******" {
			s.SecretKey = global.Config.QiNiu.SecretKey
		}
		global.Config.QiNiu = s
	case conf.Ai:
		if s.SecretKey == "******" {
			s.SecretKey = global.Config.Ai.SecretKey
		}
		global.Config.Ai = s
	}

	// 改配置文件
	core.SetConf()

	res.OkWithData("更新站点配置成功", c)
	return
}

// 动态更新前端文件

func UpdateSite(site conf.Site) error {
	if site.Project.Icon == "" && site.Project.Title == "" &&
		site.Seo.Keywords == "" && site.Seo.Description == "" &&
		site.Project.WebPath == "" {
		return nil
	} // 都为空 return 不会往后走

	if site.Project.WebPath == "" {
		return errors.New("请配置前端地址")
	}

	file, err := os.Open(site.Project.WebPath)
	if err != nil {
		return errors.New(fmt.Sprintf("%s 文件不存在", site.Project.WebPath))
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		logrus.Errorf("goquery 解析失败 %s", err)
		return errors.New("文件解析失败")
	}

	if site.Project.Title != "" {
		doc.Find("title").SetText(site.Project.Title)
	}
	if site.Project.Icon != "" {
		selection := doc.Find("link[rel=\"icon\"]")
		if selection.Length() > 0 {
			// 有就修改
			selection.SetAttr("href", site.Project.Icon)
		} else {
			// 没有就创建
			doc.Find("head").AppendHtml(fmt.Sprintf("<link rel=\"icon\" href=\"%s\">", site.Project.Icon))
		}
	}
	if site.Seo.Keywords != "" {
		selection := doc.Find("meta[name=\"keywords\"]")
		if selection.Length() > 0 {
			selection.SetAttr("content", site.Seo.Keywords)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<meta name=\"keywords\" content=\"%s\">", site.Seo.Keywords))
		}
	}
	if site.Seo.Description != "" {
		selection := doc.Find("meta[name=\"description\"]")
		if selection.Length() > 0 {
			selection.SetAttr("content", site.Seo.Description)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<meta name=\"description\" content=\"%s\">", site.Seo.Description))
		}
	}

	html, err := doc.Html()
	if err != nil {
		logrus.Errorf("生成html失败 %s", err)
		return errors.New("生成html失败")
	}

	err = os.WriteFile(site.Project.WebPath, []byte(html), 0666)
	if err != nil {
		logrus.Errorf("文件写入失败 %s", err)
		return errors.New("文件写入失败")
	}

	return nil
} // 动态更新前端文件
