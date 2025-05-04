package main

import (
	"blogW_server/core"
	"blogW_server/flags"
	"blogW_server/global"
	"blogW_server/service/redis_service/redis_jwt"
	"blogW_server/utils/jwts"
	"fmt"
)

func main() {
	flags.Parse()                   // 环境变量参数
	global.Config = core.Readconf() // 读配置文件
	core.InitLogrus()
	global.Redis = core.InitRedis()

	token, err := jwts.GetToken(jwts.Claims{
		UserID: 2,
		Role:   1,
	})
	fmt.Println(token, err)
	redis_jwt.TokenBlack(token, redis_jwt.UserBlackType)
	blk, ok := redis_jwt.HasTokenBlack(token)
	fmt.Println(blk, ok)
}
