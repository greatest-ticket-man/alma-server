package config

import (
	"github.com/BurntSushi/toml"
)

// ConfigData .
var ConfigData *AlmaConfig

// AlmaConfig .
type AlmaConfig struct {
	HTTPServer HTTPServer `toml:"http_server"`
}

// HTTPServer http serverの設定
type HTTPServer struct {
	Address     string `toml:"address"`
	AllowOrigin string `toml:"allow_origin"`
}

// Setup configファイルからデータを取得する
func Setup(path string) *AlmaConfig {

	config := &AlmaConfig{}
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		panic(err)
	}

	// set
	ConfigData = config
	return config
}
