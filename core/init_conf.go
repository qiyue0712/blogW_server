package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var confPath = "settings.yaml"

type System struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}
type Config struct {
	System System `yaml:"system"`
}

func Readconf() {
	// 读取yaml返回byte[] 解析 并映射到结构体
	byteData, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(byteData, &config)
	if err != nil {
		panic(fmt.Sprintf("yaml配置文件格式错误 %s", err))
	}
	fmt.Println(config)
}
