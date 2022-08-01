package main

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	File    *file    `mapstructure:"file"`
	Burst   *burst   `mapstructure:"burst"`
	Connect *connect `mapstructure:"connect"`
	Os      *os      `mapstructure:"os"`
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

type os struct {
	Base      []string `mapstructure:"base"`
	OpenEuler []string `mapstructure:"openEuler"`
}

var global = new(Config)

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = viper.Unmarshal(global)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
