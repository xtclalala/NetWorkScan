package main

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	File    *file    `mapstructure:"file"`
	Burst   *burst   `mapstructure:"burst"`
	Connect *connect `mapstructure:"connect"`
	Shell   *shell   `mapstructure:"shell"`
}

type file struct {
	InFileName  string `mapstructure:"inFileName"`
	OutFileName string `mapstructure:"outFileName"`
	Sheet       string `mapstructure:"sheet"`
	Ip          int    `mapstructure:"ip"`
	Port        int    `mapstructure:"port"`
	User        int    `mapstructure:"user"`
	Password    int    `mapstructure:"password"`
}

type burst struct {
	BurstNum int `mapstructure:"burstNum"`
}

type connect struct {
	Timeout int64 `mapstructure:"timeout"`
}

type shell struct {
	Name      string `mapstructure:"name"`
	AbsPath   string `mapstructure:"absPath"`
	LocalPath string `mapstructure:"localPath"`
}

var global = new(Config)

func InitConfig(configPath string) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = viper.Unmarshal(global)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
