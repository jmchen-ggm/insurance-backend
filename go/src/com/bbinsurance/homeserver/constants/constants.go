package constants

import (
	"com/bbinsurance/log"
	"github.com/kylelemons/go-gypsy/yaml"
)

var STATIC_FOLDER string
var PORT string

func InitConstants() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		log.Error("InitConstants Err %s", err)
	}
	STATIC_FOLDER, _ = config.Get("static_folder")
	PORT, _ = config.Get("port")
}
