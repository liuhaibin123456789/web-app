package tool

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var V *viper.Viper

func Viper() (err error) {
	//配置文件读取
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigFile("config.yaml")
	err = v.ReadInConfig()
	if err != nil {
		return err
	}
	//监控配置文件，动态实现配置加载
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println("某人修改了配置文件...")
	})

	V = v
	return nil
}
