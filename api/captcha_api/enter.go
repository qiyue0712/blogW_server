package captcha_api

import (
	"blogW_server/common/res"
	"blogW_server/global"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type CaptchaApi struct {
}

type CaptchaResponse struct {
	CaptchaID string `json:"captchaID"`
	Captcha   string `json:"captcha"`
}

func (CaptchaApi) CaptchaView(c *gin.Context) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      1,     // 随机噪点数量
		ShowLineOptions: 2 | 4, // 1:直线 2:曲线  4:点线
		Length:          4,
		Source:          "1234567890",
		//BgColor: &color.RGBA{
		//	R: 3,
		//	G: 102,
		//	B: 214,
		//	A: 125,
		//},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, global.CaptchaStore)
	lid, lb64s, err := captcha.Generate()
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg("图片验证码生成失败", c)
		return
	}
	res.OkWithData(CaptchaResponse{
		CaptchaID: lid,
		Captcha:   lb64s,
	}, c)
} // 图片验证码接口
