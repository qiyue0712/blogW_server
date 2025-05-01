package flags

import "flag"

type Options struct {
	File    string
	DB      bool
	Version bool
}

var FlagOptions = new(Options)

// flag函数注册参数绑定， f->FlagOptions.File 默认value
// 先解析，然后映射到FlagOptions全局使用

func Parse() {
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.BoolVar(&FlagOptions.Version, "v", false, "版本")
	flag.Parse()
}
