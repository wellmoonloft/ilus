package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/ilus/utils"
)

func init() {
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//初始化日志
	utils.InitLogs()
	//初始化缓存
	if beego.AppConfig.String("cache::is_cache") == "true" {
		utils.InitCache()
	}
	//初始化数据库
	InitDatabase()
}
