package configs

import (
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	DB database `toml:"database"`
}

type database struct {
	Server   string
	Ports    string
	ConnMax  int `toml:"connection_max"`
	Enabled  bool
	UserName string `toml:"username"`
	Passwd   string
	DBName   string `toml:"dbname"`
}

var config tomlConfig

func init() {
	filePath := "../../configs/config.toml"
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}
}

func GetConfig() tomlConfig {
	return config
}
