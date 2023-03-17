package cfg

import cfg "github.com/GiHccTpD/go-kit/config"

func Init() {
	// 加载配置
	cfg.LoadConfigFile("", "./conf.template.toml")
}
