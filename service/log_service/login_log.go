package log_service

import (
	"blogW_server/core"
	"blogW_server/global"
	"blogW_server/models"
	"blogW_server/models/enum"
	"blogW_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

// 成功的登录日志

func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	claims, err := jwts.ParseTokenByGin(c) // 调用获取token的函数获取userID
	userID := uint(0)
	userName := ""
	if err == nil && claims != nil {
		userID = claims.UserID
		userName = claims.UserName
	}

	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录",
		Content:     "",
		UserID:      userID,
		IP:          ip,
		Addr:        addr,
		LoginStatus: true,
		Username:    userName,
		Pwd:         "",
		LoginType:   loginType,
	})
}

// 失败的登录日志

func NewLoginFail(c *gin.Context, loginType enum.LoginType, msg string, username string, pwd string) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录失败",
		Content:     msg,
		IP:          ip,
		Addr:        addr,
		LoginStatus: false,
		Username:    username,
		Pwd:         pwd,
		LoginType:   loginType,
	})
}
