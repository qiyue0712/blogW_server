package core

import (
	"blogW_server/conf"
	"blogW_server/flags"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func Readconf() (c *conf.Config) {
	// 读取yaml返回byte[] 解析 并映射到结构体config
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}
	c = new(conf.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		panic(fmt.Sprintf("yaml配置文件格式错误 %s", err))
	}
	fmt.Printf("读取配置文件 %s 成功\n", flags.FlagOptions.File)
	return
}
