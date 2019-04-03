package sysinit

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SysInfo struct {
	Language        string
	Languagecode    string
	Systitle        string
	Sysurl          string
	Logo            string
	Favicon         string
	Keyword         []string
	Abs             string
	Baidutoken      string
	Googlecaptcha   string
	Baiducaptcha    string
	Bingcaptcha     string
	Indexnum        int
	Rssnum          int
	Absnum          int
	Aevterid        int
	Aevtername      string
	Aevterurl       string
	Isexamine       bool
	IsnewInform     bool
	Isexamineinform bool
	Isreplyinform   bool
	Iscommentapi    bool
	Commentnum      int
	Position        int
	Positionurl     string
	Aaccesskey      string
	Secretkey       string
	Bucket          string
	Strategy        string
	Ispjax          bool
	Layout          int
	Issidebar       bool
	Isemail         bool
	Smtp            string
	Emailname       string
	Emailpwd        string
	Sendname        string
	Isaip           bool
	Token           string
}

func Yamlmain() {

	c := SysInfo{}
	yamlFile, err := ioutil.ReadFile("./conf/sysInfo.yml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	//return c
	fmt.Println(c)
}
