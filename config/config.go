package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = viper.New()

func init() {
	//加载配置文件
	Conf.SetConfigFile("./config/config.yaml")

	if err := Conf.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	go reloadConfig()
}

// 热加载
func reloadConfig() {
	Conf.WatchConfig()
	Conf.OnConfigChange(func(event fsnotify.Event) {
		fmt.Printf("Detect config change: %s \n", event.String())
	})
}
