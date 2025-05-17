package user_api

import (
	"blogW_server/common/res"
	"blogW_server/global"
	"blogW_server/middleware"
	"blogW_server/models"
	"blogW_server/service/user_service"
	"blogW_server/utils/jwts"
	"blogW_server/utils/pwd"
	"github.com/gin-gonic/gin"
)

type PwdLoginRequest struct {
	Val      string `json:"val" binding:"required"` // 有可能是用户名, 也有可能是邮箱
	Password string `json:"password" binding:"required"`
}

func (UserApi) PwdLoginApi(c *gin.Context) {

	cr := middleware.GetBind[PwdLoginRequest](c)

	if !global.Config.Site.Login.UsernamePwdLogin {
		res.FailWithMsg("站点未启用密码登录", c)
	}

	var user models.UserModel
	err := global.DB.Take(&user, "(username = ? or email = ?) and password <> ''",
		cr.Val, cr.Val).Error
	if err != nil {
		res.FailWithMsg("用户名密码错误", c)
		return
	}
	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		res.FailWithMsg("用户名密码错误", c)
		return
	}

	// 颁发token
	token, _ := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		UserName: user.Username,
		Role:     user.Role,
	})
	user_service.NewUserService(user).UserLogin(c)

	res.OkWithData(token, c)
} // 用户管理--用户密码登录
