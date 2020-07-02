package config

import (
	"github.com/BurntSushi/toml"
)

// ConfigData .
var ConfigData *AlmaConfig

// AlmaConfig .
type AlmaConfig struct {
	HTTPServer     *HTTPServer `toml:"httpserver"`
	MongoDatabases []*MongoDB  `toml:"mongodatabases"`
}

// HTTPServer http serverの設定
type HTTPServer struct {
	Address     string `toml:"address"`
	AllowOrigin string `toml:"allow_origin"`
}

// MongoDB MongoDB設定
type MongoDB struct {
	Host     string         `toml:"host"`
	Port     string         `toml:"port"`
	Db       string         `toml:"db"`
	User     string         `toml:"user"`
	Password string         `toml:"password"`
	Option   *MongoDBOption `toml:"option"`
}

// MongoDBOption MongoDB option
type MongoDBOption struct {
	ConnectTimeoutMs    string `toml:"connecttimeoutms"`
	HeartBeatIntervalMs string `toml:"heartbeatintervalms"`
	MaxIdleTimeMs       string `toml:"maxidletimems"`
	MaxPoolSize         string `toml:"maxpoolsize"`
	ReadPreference      string `toml:"readpreference"`
	ReadConcernLevel    string `toml:"readconcernlevel"`
	WriteConnection     string `toml:"w"`
	SocketTimeoutMs     string `toml:"sockettimeoutms"`
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
