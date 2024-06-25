package market

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()
	logx.MustSetup()
}
