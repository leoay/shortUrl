package configs

import (
	"time"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
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

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

var config tomlConfig

func init() {
	filePath := "../../configs/config.toml"
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}
}

func GetVConfig() tomlConfig {
	return config
}
