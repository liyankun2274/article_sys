package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Port               string `toml:"port"`
	MysqlDb            string `toml:"mysql_db"`
	RedisUrl           string `toml:"redis_url"`
	RedisSessionPrefix string `toml:"redis_session_prefix"`
}

func LoadConfig(c string) (*Config, error) {
	var config Config
	var _,err = toml.DecodeFile(c,&config)
	if err != nil{
		return nil,err
	}else {
		return &config,nil
	}
}
