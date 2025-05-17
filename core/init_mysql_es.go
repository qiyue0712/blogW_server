package core

import (
	"blogW_server/global"
	river "blogW_server/service/river_service"
	"github.com/sirupsen/logrus"
)

func InitMysqlES() {
	if !global.Config.River.Enable {
		logrus.Infof("关闭mysql同步操作")
		return
	}

	r, err := river.NewRiver()
	if err != nil {
		logrus.Error(err)
	}
	go r.Run()
}
