package qiniu_service

import (
	"blogW_server/global"
	"context"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
	"time"
)

func GetToken() (token string, err error) {
	mac := credentials.NewCredentials(global.Config.QiNiu.AccessKey, global.Config.QiNiu.SecretKey)
	putPolicy, err := uptoken.NewPutPolicy(global.Config.QiNiu.Bucket, time.Now().Add(time.Duration(global.Config.QiNiu.Expiry)*time.Second))
	if err != nil {
		return
	}
	token, err = uptoken.NewSigner(putPolicy, mac).GetUpToken(context.Background())
	if err != nil {
		return
	}
	return
}
