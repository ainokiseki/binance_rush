package settings

import (
	"log"

	"github.com/go-ini/ini"
)

type BinanceConfig struct {
	AK string `ini:"ak"`
	SK string `ini:"sk"`
}

var Config BinanceConfig

func init() {
	cfg, err := ini.Load("./settings/binance.ini")
	if err != nil {
		log.Fatal("read ak/sk fail:", err)
	}

	err = cfg.MapTo(&Config)
	if err != nil {
		log.Fatal("map ak/sk fail:", err)
	}
}
