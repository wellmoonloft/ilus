package main

import (
	"github.com/astaxie/beego"
	_ "github.com/ilus/routers"
	_ "github.com/ilus/sysinit"
)

func main() {
	beego.Run()
}
