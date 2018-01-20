package constants

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

var STATIC_FOLDER string
var LOGIC_DB_PATH string

func InitConstants() {
	config, _ := yaml.ReadFile("conf.yaml")
	STATIC_FOLDER, _ = config.Get("static_folder")
	LOGIC_DB_PATH, _ = config.Get("logic_db_path")
}
