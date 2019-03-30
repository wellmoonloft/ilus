package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/ilus/models"
	"github.com/ilus/utils"
	"strings"
)

var cpt *captcha.Captcha

//HomeController 首页
type HomeController struct {
	BaseController
}

/*func init() {
	//来个验证码
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}*/

//Index 首页
func (c *HomeController) Index() {
	//判断是否登录
	c.checkLogin()
	c.setTpl()
	c.Data["pageTitle"] = beego.AppConfig.String("appname") + " | 首页"
}

//Page404 404
func (c *HomeController) Page404() {
	c.setTpl()
}

//Error 其他错误
func (c *HomeController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.setTpl("home/error.html", "shared/layout_pullbox.html")
}

//Login 登陆页面
func (c *HomeController) Login() {

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "home/login_headcssjs.html"
	c.LayoutSections["footerjs"] = "home/login_footerjs.html"
	c.Data["pageTitle"] = beego.AppConfig.String("appname") + " | 后台登陆"
	c.setTpl("home/login.html", "shared/layout_base.html")
}

//DoLogin 登陆操作
func (c *HomeController) DoLogin() {

	username := strings.TrimSpace(c.GetString("UserName"))
	userpwd := strings.TrimSpace(c.GetString("UserPwd"))
	/*captchaID := strings.TrimSpace(c.GetString("captcha_id"))
	Captcha := strings.TrimSpace(c.GetString("Captcha"))
	if !cpt.Verify(captchaID, Captcha) {
		c.jsonResult(utils.JRCodeFailed, "验证码不正确", "")
	}*/
	if len(username) == 0 || len(userpwd) == 0 {
		c.jsonResult(utils.JRCodeFailed, "用户名和密码不正确", "")
	}
	userpwd = utils.String2md5(userpwd)
	user, err := models.BackendUserOneByUserName(username, userpwd)
	if user != nil && err == nil {
		if user.Status == utils.Disabled {
			c.jsonResult(utils.JRCodeFailed, "用户被禁用，请联系管理员", "")
		}
		//保存用户信息到session
		c.setBackendUser2Session(user.Id)
		//获取用户信息
		c.jsonResult(utils.JRCodeSucc, "登录成功", "")
	} else {
		c.jsonResult(utils.JRCodeFailed, "用户名或者密码错误", "")
	}
	//utils.LogInfo(username + "登陆成功" + time.Now().String())
}

//Logout 登出操作
func (c *HomeController) Logout() {
	user := models.BackendUser{}
	c.SetSession("backenduser", user)
	c.pageLogin()
	//utils.LogInfo(strings.TrimSpace(c.GetString("UserName")) + "已登出" + time.Now().String())
}
