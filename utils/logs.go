package utils

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger

// fileLogs 生产环境下日志
var fileLogs *logs.BeeLogger

//运行方式
var runmode string

//InitLogs 初始化日志配置
func InitLogs() {
	consoleLogs = logs.NewLogger(1)
	consoleLogs.SetLogger(logs.AdapterConsole)
	consoleLogs.Async() //异步
	fileLogs = logs.NewLogger(10000)
	level := beego.AppConfig.String("logs::level")
	fileLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/ilus.log",
		"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
		"level":`+level+`,
		"daily":true,
		"maxdays":10}`)
	fileLogs.Async() //异步
	runmode = strings.TrimSpace(strings.ToLower(beego.AppConfig.String("runmode")))
	if runmode == "" {
		runmode = "dev"
	}
}

//LogEmergency 紧急
func LogEmergency(v interface{}) {
	format := "%s"
	if runmode == "debug" {
		fileLogs.Emergency(format, v)
	}
	consoleLogs.Emergency(format, v)
}

//LogAlert 警报
func LogAlert(v interface{}) {
	format := "%s"
	if runmode == "debug" {
		fileLogs.Alert(format, v)
	}
	consoleLogs.Alert(format, v)
}

//LogCritical 严重
func LogCritical(v interface{}) {
	format := "%s"
	if runmode == "debug" {
		fileLogs.Critical(format, v)
	}
	consoleLogs.Critical(format, v)
}

//LogError 错误
func LogError(v interface{}) {
	format := "%s"
	if runmode == "debug" {
		fileLogs.Error(format, v)
	}
	consoleLogs.Error(format, v)
}

//LogWarning 警告
func LogWarning(v interface{}) {
	format := "%s"
	if runmode == "debug" {
		fileLogs.Warning(format, v)
	}
	consoleLogs.Warning(format, v)
}

//LogNotice 通知
func LogNotice(v interface{}) {
	format := "%s"
	if runmode == "debug" {
		fileLogs.Notice(format, v)
	}
	consoleLogs.Notice(format, v)
}

//LogInfo 信息
func LogInfo(v interface{}) {
	format := "%s"
	if runmode == "debug" {
		fileLogs.Info(format, v)
	}
	consoleLogs.Info(format, v)
}

//LogDebug 调试
func LogDebug(v interface{}) {
	format := "%s"
	if runmode == "debug" {
		fileLogs.Debug(format, v)
	}
	consoleLogs.Debug(format, v)
}

//LogTrace 运行方式
func LogTrace(v interface{}) {
	format := "%s"
	if runmode == "debug" {
		fileLogs.Trace(format, v)
	}
	consoleLogs.Trace(format, v)
}
