package user_api

import (
	"blogW_server/common/res"
	"blogW_server/global"
	"blogW_server/models"
	"blogW_server/models/enum"
	"blogW_server/service/qq_service"
	"blogW_server/service/user_service"
	"blogW_server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type QQLoginRequest struct {
	Code string `json:"code" binding:"required"`
}

func (UserApi) QQLoginView(c *gin.Context) {
	var cr QQLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	if !global.Config.Site.Login.UsernamePwdLogin {
		res.FailWithMsg("站点未启用qq登录", c)
		return
	}
	info, err := qq_service.GetUserInfo(cr.Code)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, "open_id = ?", info.OpenID).Error
	if err != nil {
		// 创建用户
		uname := base64Captcha.RandText(5, "0123456789")
		user = models.UserModel{
			Username:       fmt.Sprintf("b_%s", uname),
			Nickname:       info.Nickname,
			Avatar:         info.Avatar,
			RegisterSource: enum.RegisterQQSourceType,
			OpenID:         info.OpenID,
			Role:           enum.UserRole,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			logrus.Error(err)
			res.FailWithMsg("qq登录失败", c)
			return
		}
	}

	// 颁发token
	token, _ := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		UserName: user.Username,
		Role:     user.Role,
	})
	user_service.NewUserService(user).UserLogin(c)
	res.OkWithData(token, c)

}
