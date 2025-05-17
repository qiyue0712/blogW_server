package api

import (
	"blogW_server/api/banner_api"
	"blogW_server/api/captcha_api"
	"blogW_server/api/image_api"
	"blogW_server/api/log_api"
	"blogW_server/api/site_api"
	"blogW_server/api/user_api"
)

type Api struct {
	SiteApi    site_api.SiteApi
	LogApi     log_api.LogApi
	ImageApi   image_api.ImageApi
	BannerApi  banner_api.BannerApi
	CaptchaApi captcha_api.CaptchaApi
	UserApi    user_api.UserApi
}

var App = Api{} // 自动初始化子模块结构体
