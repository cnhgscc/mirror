package cmdargs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	pflag.StringP("server.config", "c", "./config.yaml", "server.config")
	pflag.StringP("server.host", "h", "0.0.0.0", "server.host")
	pflag.IntP("server.port", "p", 9000, "server.port")
	pflag.StringP("server.name", "n", "grpc", "server.name")
}

type Option func()

func Init(opts ...Option) {
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigFile(viper.GetString("server.config"))
	_ = viper.ReadInConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		for _, opt := range opts {
			opt()
		}
	})
	viper.WatchConfig()
}
