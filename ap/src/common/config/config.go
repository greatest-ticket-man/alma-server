package config

import (
	"alma-server/ap/src/common/util/cryptoutil"
	"strings"

	"github.com/BurntSushi/toml"
)

// ConfigData .
var ConfigData *AlmaConfig

// AlmaConfig .
type AlmaConfig struct {
	Mode             string      `toml:"mode"`
	RootDirectory    string      `toml:"rootdirectory"`
	MasterCacheDir   string      `toml:"mastercachedir"`
	HTTPServer       *HTTPServer `toml:"httpserver"`
	PrometheusServer *HTTPServer `toml:"prometheus"`
	Mail             *Mail       `toml:"mail"`
	Stripe           *Stripe     `toml:"stripe"`
	MongoDatabases   []*MongoDB  `toml:"mongodatabases"`
}

// HTTPServer http serverの設定
type HTTPServer struct {
	Address     string `toml:"address"`
	AllowOrigin string `toml:"allow_origin"`
	TLS         bool   `toml:"tls"`
	CertFile    string `toml:"certfile"`
	KeyFile     string `toml:"keyfile"`
}

// Mail メールの設定
type Mail struct {
	Gmail *Gmail `toml:"gmail"`
}

// Gmail Gmailの設定
type Gmail struct {
	Address  string `toml:"address"`
	Password string `toml:"password"`
}

// Stripe stripe cledit clientの設定
type Stripe struct {
	PublicKey string `toml:"publickey"`
	SecretKey string `toml:"secretkey"`
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

	// パスワードの暗号化解除
	if config.Mail != nil {
		if config.Mail.Gmail != nil {
			config.Mail.Gmail.Password = strings.TrimSpace(cryptoutil.DecPassword(config.Mail.Gmail.Password))
		}
	}

	// set
	ConfigData = config
	return config
}
