package main

import (
	"blogW_server/core"
	"fmt"
)

func main() {
	core.InitIPDB()
	add := core.GetIpAddr("175.0.201.207")
	fmt.Println(add)
}
