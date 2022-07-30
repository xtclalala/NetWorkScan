package main

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	File *file `mapstructure:"file"`
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
