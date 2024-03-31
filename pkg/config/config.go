package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *viper.Viper

// var ConfigMap map[string]interface{}
func Start() {
	ConfigJson()
}

func ConfigJson() {
	Config = viper.New()
	Config.AddConfigPath("./pkg/config")
	Config.SetConfigName("config.json")
	Config.SetConfigType("json")
	if err := Config.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
