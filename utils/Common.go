package utils

//返回格式
type JsonResultCode int

//常量
const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至地址
	JRCode401 = 401 //未授权访问
)

//常量
const (
	Deleted = iota - 1
	Disabled
	Enabled
)
