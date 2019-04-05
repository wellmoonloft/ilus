package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ilus/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
)

//SystemController 分类管理
type SystemController struct {
	BaseController
}

//SystemSettings 系统配置
type SystemSettings struct {
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
	Isapi              bool
	Token              string
}

//Index 系统设置首页
func (c *SystemController) Index() {
	c.Data["pageTitle"] = beego.AppConfig.String("appname") + " | 系统设定"
	c.Data["m"] = c.SystemSettings
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "system/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "system/index_footerjs.html"
}

//getSystemInfo 初始化系统设置
func (s *SystemSettings) getSystemInfo() {

	yamlFile, err := ioutil.ReadFile("./conf/sysInfo.yml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, s)
	if err != nil {
		fmt.Println(err.Error())
	}
	//return c
	//fmt.Println(y)
}

//Save 保存当前系统设置
func (c *SystemController) Save() {

	action := c.GetString(":action")
	switch action {
	case "general":

		//c.SystemSettings.Languagecode = c.Ctx.Request.PostForm.Get("blog_locale")
		c.SystemSettings.Systitle = c.Ctx.Request.PostForm.Get("blog_title")
		c.SystemSettings.Sysurl = c.Ctx.Request.PostForm.Get("blog_url")
		c.SystemSettings.Logo = c.Ctx.Request.PostForm.Get("blog_logo")
		c.SystemSettings.Favicon = c.Ctx.Request.PostForm.Get("blog_favicon")

	case "seo":

		//c.SystemSettings.Keyword[0] = c.Ctx.Request.PostForm.Get("seo_keywords")
		c.SystemSettings.Abs = c.Ctx.Request.PostForm.Get("seo_desc")
		c.SystemSettings.Baidutoken = c.Ctx.Request.PostForm.Get("seo_baidu_token")
		c.SystemSettings.Googlecaptcha = c.Ctx.Request.PostForm.Get("blog_verification_google")
		c.SystemSettings.Bingcaptcha = c.Ctx.Request.PostForm.Get("blog_verification_bing")
		c.SystemSettings.Baiducaptcha = c.Ctx.Request.PostForm.Get("blog_verification_baidu")

	case "article":

		i, _ := strconv.Atoi(c.Ctx.Request.PostForm.Get("index_posts"))
		c.SystemSettings.Indexnum = i
		i, _ = strconv.Atoi(c.Ctx.Request.PostForm.Get("rss_posts"))
		c.SystemSettings.Rssnum = i
		i, _ = strconv.Atoi(c.Ctx.Request.PostForm.Get("post_summary"))
		c.SystemSettings.Absnum = i

	case "comment":

		c.SystemSettings.Aevterurl = c.Ctx.Request.PostForm.Get("native_comment_avatar")
		b, _ := strconv.ParseBool(c.Ctx.Request.PostForm.Get("new_comment_need_check"))
		c.SystemSettings.Isexamine = b
		b, _ = strconv.ParseBool(c.Ctx.Request.PostForm.Get("new_comment_notice"))
		c.SystemSettings.IsnewInform = b
		b, _ = strconv.ParseBool(c.Ctx.Request.PostForm.Get("comment_pass_notice"))
		c.SystemSettings.Isexamineinform = b
		b, _ = strconv.ParseBool(c.Ctx.Request.PostForm.Get("comment_reply_notice"))
		c.SystemSettings.Isreplyinform = b
		b, _ = strconv.ParseBool(c.Ctx.Request.PostForm.Get("comment_api_switch"))
		c.SystemSettings.Iscommentapi = b
		i, _ := strconv.Atoi(c.Ctx.Request.PostForm.Get("index_comments"))
		c.SystemSettings.Commentnum = i
		c.SystemSettings.CommentPlaceholder = c.Ctx.Request.PostForm.Get("native_comment_placeholder")

	case "attach":

		i, _ := strconv.Atoi(c.Ctx.Request.PostForm.Get("attach_loc"))
		c.SystemSettings.Position = i
		c.SystemSettings.Positionurl = c.Ctx.Request.PostForm.Get("qiniu_domain")
		c.SystemSettings.Accesskey = c.Ctx.Request.PostForm.Get("qiniu_access_key")
		c.SystemSettings.Secretkey = c.Ctx.Request.PostForm.Get("qiniu_secret_key")
		c.SystemSettings.Bucket = c.Ctx.Request.PostForm.Get("qiniu_bucket")
		c.SystemSettings.Strategy = c.Ctx.Request.PostForm.Get("qiniu_small_url")

	case "admin":

		b, _ := strconv.ParseBool(c.Ctx.Request.PostForm.Get("admin_pjax"))
		c.SystemSettings.Ispjax = b
		b, _ = strconv.ParseBool(c.Ctx.Request.PostForm.Get("sidebar_style"))
		c.SystemSettings.Issidebar = b

	case "email":

		b, _ := strconv.ParseBool(c.Ctx.Request.PostForm.Get("smtp_email_enable"))
		c.SystemSettings.Isemail = b
		c.SystemSettings.Smtp = c.Ctx.Request.PostForm.Get("mail_smtp_host")
		c.SystemSettings.Emailname = c.Ctx.Request.PostForm.Get("emailSmtpUserName")
		c.SystemSettings.Emailpwd = c.Ctx.Request.PostForm.Get("emailSmtpPassword")
		c.SystemSettings.Sendname = c.Ctx.Request.PostForm.Get("mail_from_name")

	case "other":

		b, _ := strconv.ParseBool(c.Ctx.Request.PostForm.Get("api_status"))
		c.SystemSettings.Isapi = b
		c.SystemSettings.Token = c.Ctx.Request.PostForm.Get("api_token")

	default:
		fmt.Println("Unknown!")
	}

	d, err := yaml.Marshal(c.SystemSettings)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("./conf/sysInfo.yml", d, os.ModeAppend)

	c.jsonResult(utils.JRCodeSucc, "保存成功", 0)

}
