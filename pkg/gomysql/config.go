package gomysql

type Config struct {
	Scope string `yaml:"scope"`

	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	DB   string `yaml:"db"`
	Addr string `yaml:"addr"`
	Max  int    `yaml:"max"`
}
