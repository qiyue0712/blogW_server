package main

import (
	"fmt"
	"reflect"
)

type UserInfoUpdateRequest struct {
	Username    *string   `json:"username" s-u:"username"`
	Nickname    *string   `json:"nickname" s-u:"nickname"`
	Avatar      *string   `json:"avatar" s-u:"avatar"`
	Abstract    *string   `json:"abstract" s-u:"abstract"`
	LikeTags    *[]string `json:"likeTags" s-u-c:"like_tags"`
	OpenCollect *bool     `json:"openCollect" s-u-c:"open_collect"` // 公开我的收藏
	OpenFollow  *bool     `json:"openFollow" s-u-c:"open_follow"`   // 公开我的关注
	OpenFans    *bool     `json:"openFans" s-u-c:"open_fans"`       // 公开我的粉丝
	HomeStyleID *uint     `json:"homeStyleID" s-u-c:"homeStyleID"`  // 主页样式的id
}

func StructToMap(data any, t string) (mp map[string]any) {
	mp = make(map[string]any)
	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i)
		tag := v.Type().Field(i).Tag.Get(t)
		if tag == "" || tag == "-" {
			continue
		}
		if val.IsNil() {
			continue
		}
		if val.Kind() == reflect.Ptr {
			mp[tag] = val.Elem().Interface()
			continue
		}
		mp[tag] = val.Interface()
	}
	return
}

func main() {
	var name = "wangwinter"
	var openFans = true
	var cr = UserInfoUpdateRequest{
		Nickname: &name,
		OpenFans: &openFans,
	}
	fmt.Println(StructToMap(cr, "s-u"))
	fmt.Println(StructToMap(cr, "s-u-c"))
}
