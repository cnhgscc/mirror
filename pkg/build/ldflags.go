package build

import (
	"github.com/spf13/viper"
)

// ldflags args

var (
	Name    string //  项目名称
	Version string // 构建时 git 的版本
)

func CMDName() string {
	if Name == "" {
		return viper.GetString("server.name")
	}
	return Name
}
