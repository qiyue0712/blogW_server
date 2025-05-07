package file

import (
	"blogW_server/global"
	"blogW_server/utils"
	"errors"
	"strings"
)

func ImageSuffixJudge(filename string) (suffix string, err error) {
	_list := strings.Split(filename, ".")
	if len(_list) == 1 {
		err = errors.New("错误的文件名")
		return
	}
	// xxx.jpg xxx xxx.jpg.exe
	suffix = _list[len(_list)-1]
	if !utils.InList(suffix, global.Config.Upload.WhiteList) {
		err = errors.New("文件非法")
		return
	}
	return
}
