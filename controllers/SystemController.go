package controllers

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

//SystemController 系统配置
type SystemController struct {
	Language           string
	Languagecode       string
	Systitle           string
	Sysurl             string
	Logo               string
	Favicon            string
	Keyword            []string
	Abs                string
	Baidutoken         string
	Googlecaptcha      string
	Baiducaptcha       string
	Bingcaptcha        string
	Indexnum           int
	Rssnum             int
	Absnum             int
	Aevterid           int
	Aevtername         string
	Aevterurl          string
	Isexamine          bool
	IsnewInform        bool
	Isexamineinform    bool
	Isreplyinform      bool
	Iscommentapi       bool
	Commentnum         int
	CommentPlaceholder string
	Position           int
	Positionurl        string
	Accesskey          string
	Secretkey          string
	Bucket             string
	Strategy           string
	Ispjax             bool
	Isayout            int
	Issidebar          bool
	Isemail            bool
	Smtp               string
	Emailname          string
	Emailpwd           string
	Sendname           string
	Isaip              bool
	Token              string
}

//getSystemInfo 初始化系统设置
func (y *SystemController) getSystemInfo() {

	yamlFile, err := ioutil.ReadFile("./conf/sysInfo.yml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, y)
	if err != nil {
		fmt.Println(err.Error())
	}
	//return c
	//fmt.Println(y)
}

//saveSystemInfo 保存当前系统设置
func (y *SystemController) saveSystemInfo() {

	d, err := yaml.Marshal(y)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("./conf/sysInfo.yml", d, os.ModeAppend)
	//return c
	//fmt.Println(y)
}
