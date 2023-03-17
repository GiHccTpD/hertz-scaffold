package lookup

import (
	"git.ccwork.com/ccwork/go/utils.git/discovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
)

var discover discovery.Discover

func Init() {
	var (
		consulAddr = viper.GetString("server.consulAddress")
	)
	_discover, err := discovery.New(discovery.WithConsul(consulAddr))
	if err != nil {
		hlog.Errorf("discovery.New failed, msg: %v", err)
		panic(err)
	}
	discover = _discover
}

func GetDiscover() discovery.Discover {
	return discover
}
