package main

import (
	"errors"
	"fmt"
	e "github.com/pkg/errors"
	"runtime"
)

func main() {
	err := errors.New("错误")
	SetError(err)
}

func SetError1(err error) {
	msg := e.WithStack(err)
	fmt.Printf("%+v\n", msg)
}

func SetError(err error) {
	var msg = make([]byte, 1024)
	n := runtime.Stack(msg, true)
	fmt.Printf(string(msg[:n]))
}
