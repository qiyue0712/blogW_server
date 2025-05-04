package conf

import "fmt"

type System struct {
	Ip      string `yaml:"ip"`
	Port    int    `yaml:"port"`
	GinMode string `yaml:"gin_mode"`
	Env     string `yaml:"env"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
