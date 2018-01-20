package constants

import (
	"com/bbinsurance/log"
	"github.com/kylelemons/go-gypsy/yaml"
)

var PORT string

func InitConstants() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		log.Error("InitConstants Err %s", err)
	}
	PORT, _ = config.Get("port")
}
