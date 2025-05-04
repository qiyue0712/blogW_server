package core

import (
	"blogW_server/global"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"log"
)

func InitRedis() *redis.Client {
	r := global.Config.Redis

	redisDB := redis.NewClient(&redis.Options{
		Addr:     r.Addr,     // 不写默认就是这个
		Password: r.Password, // 密码
		DB:       r.DB,       // 默认是0
	}) // 初始化redis
	_, err := redisDB.Ping().Result()
	if err != nil {
		log.Fatalf("redis连接失败 %s", err)
	}
	logrus.Info("redis连接成功")
	return redisDB
}
