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

//紧急
func LogEmergency(v interface{}) {
	log("emergency", v)
}

//警报
func LogAlert(v interface{}) {
	log("alert", v)
}

//严重
func LogCritical(v interface{}) {
	log("critical", v)
}

//错误
func LogError(v interface{}) {
	log("error", v)
}

//警告
func LogWarning(v interface{}) {
	log("warning", v)
}

//通知
func LogNotice(v interface{}) {
	log("notice", v)
}

//信息
func LogInfo(v interface{}) {
	log("info", v)
}

//调试
func LogDebug(v interface{}) {
	log("debug", v)
}

//运行方式
func LogTrace(v interface{}) {
	log("trace", v)
}

//Log 输出日志
func log(level, v interface{}) {
	format := "%s"
	if level == "" {
		level = "debug"
	}
	if runmode == "dev" {
		switch level {
		case "emergency":
			fileLogs.Emergency(format, v)
		case "alert":
			fileLogs.Alert(format, v)
		case "critical":
			fileLogs.Critical(format, v)
		case "error":
			fileLogs.Error(format, v)
		case "warning":
			fileLogs.Warning(format, v)
		case "notice":
			fileLogs.Notice(format, v)
		case "info":
			fileLogs.Info(format, v)
		case "debug":
			fileLogs.Debug(format, v)
		case "trace":
			fileLogs.Trace(format, v)
		default:
			fileLogs.Debug(format, v)
		}
	} else {
		switch level {
		case "emergency":
			consoleLogs.Emergency(format, v)
		case "alert":
			consoleLogs.Alert(format, v)
		case "critical":
			consoleLogs.Critical(format, v)
		case "error":
			consoleLogs.Error(format, v)
		case "warning":
			consoleLogs.Warning(format, v)
		case "notice":
			consoleLogs.Notice(format, v)
		case "info":
			consoleLogs.Info(format, v)
		case "debug":
			consoleLogs.Debug(format, v)
		case "trace":
			consoleLogs.Trace(format, v)
		default:
			consoleLogs.Debug(format, v)
		}
	}
}
