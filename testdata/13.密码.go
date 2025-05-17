package main

import (
	"blogW_server/utils/pwd"
	"fmt"
)

func main() {
	p, _ := pwd.GenerateFromPassword("123456")
	fmt.Println(p)

	fmt.Println(pwd.CompareHashAndPassword(p, "123456"))

}
