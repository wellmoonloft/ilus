package controllers

import "github.com/astaxie/beego"

type BackendSettingsController struct {
	BaseController
}

func (c *BackendSettingsController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	c.checkLogin()
}

func (c *BackendSettingsController) Index() {
	c.Data["pageTitle"] = beego.AppConfig.String("appname") + " | 系统设定"
	//需要权限控制
	c.checkAuthor()
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "backendsettings/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "backendsettings/index_footerjs.html"
}
