package main

import (
	"blogW_server/core"
	"blogW_server/flags"
	"blogW_server/global"
	"blogW_server/service/qiniu_service"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.Readconf()
	core.InitLogrus()
	//url, err := SendFile("uploads/images001/d2dfb84660362d463b3d8b356dbe67dc.jpeg")
	//fmt.Println(url, err)
	//file, _ := os.Open("uploads/images001/d2dfb84660362d463b3d8b356dbe67dc.jpeg")

	//url, err := SendReader(file)
	//fmt.Println(url, err)
	fmt.Println(qiniu_service.GetToken())
}
