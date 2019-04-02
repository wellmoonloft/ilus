package sysinit

import (
	"fmt"
	"github.com/astaxie/beego/config/yaml"
)

type CSDN struct {
	Language     string
	LanguageCode string
	SysTitle     string
	SysUrl       string
	Logo         string
	favicon      string
	KeyWord      []string
}

func Yamlmain() {
	conf, err := yaml.ReadYmlReader("./conf/sysInfo.yml")
	if err != nil {
		fmt.Println(err)
	}
	name := conf["languageCode"]
	favourite := conf["favicon"]
	fmt.Println(conf["language"])
	fmt.Println(conf["baiduToken"])
	fmt.Println(favourite)
	fmt.Println(name)
}
