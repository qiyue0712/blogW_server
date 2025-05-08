package main

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var stores = base64Captcha.DefaultMemStore

func main() {
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
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		//Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, stores)
	lid, lb64s, lerr := captcha.Generate()
	code := stores.Get(lid, true)
	fmt.Println(lid)
	fmt.Println(lb64s)
	fmt.Println(lerr)
	fmt.Println(code)
}
