package settings

import (
	"log"
	"web_app/dao/mysql"
	"web_app/tool"
)

func Settings() {
	//加载配置
	if err := tool.Viper(); err != nil {
		log.Fatal("viper出错:", err)
		return
	}
	//初始化日志
	if err := tool.Logger(); err != nil {
		log.Fatal("zap出错:", err)
		return
	}
	tool.SugaredDebug("zap logger初始化...")

	if err := mysql.Mysql(); err != nil {
		tool.SugaredPanicf("mysql init error: %s", err.Error())
		return
	}
	tool.SugaredDebug("mysql 初始化...")

	if err := tool.Redis(); err != nil {
		tool.SugaredPanicf("redis init error: %s", err.Error())
		return
	}
	tool.SugaredDebug("redis 初始化...")

	if err := tool.InitSnowflake(); err != nil {
		tool.SugaredPanicf("snowflake init error: %s", err.Error())
		return
	}
	tool.SugaredDebug("snowflake 节点初始化...")
}
