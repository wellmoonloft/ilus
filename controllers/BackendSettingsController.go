package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ilus/models"
	"github.com/ilus/utils"
)

//BackendSettingsController 系统设置
type BackendSettingsController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *BackendSettingsController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	c.checkLogin()
}

//Index 系统设置首页
func (c *BackendSettingsController) Index() {
	c.Data["pageTitle"] = beego.AppConfig.String("appname") + " | 系统设定"
	m, err := models.BackendSettingsOne(1)
	if m == nil || err != nil {
		c.pageError("数据无效，请刷新后重试")
	}
	c.Data["m"] = m
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "backendsettings/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "backendsettings/index_footerjs.html"
}

//GeneralSave 基础信息保存
func (c *BackendSettingsController) GeneralSave() {
	oM, err := models.BackendSettingsOne(1)
	if oM == nil || err != nil {
		c.jsonResult(utils.JRCodeFailed, "数据无效，请刷新后重试", "")
	}
	m := models.BackendSettings{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(utils.JRCodeFailed, "获取数据失败", m.Id)
	}
	/*oM.RealName = m.RealName
	oM.Mobile = m.Mobile
	oM.Email = m.Email
	oM.Avatar = c.GetString("ImageUrl")*/
	o := orm.NewOrm()
	if _, err := o.Update(oM); err != nil {
		c.jsonResult(utils.JRCodeFailed, "编辑失败", m.Id)
	} else {
		c.jsonResult(utils.JRCodeSucc, "保存成功", m.Id)
	}
}

//PasswordSave 密码保存
func (c *BackendSettingsController) PasswordSave() {
	Id := c.curUser.Id
	oM, err := models.BackendSettingsOne(Id)
	if oM == nil || err != nil {
		c.pageError("数据无效，请刷新后重试")
	}
	/*oldPwd := strings.TrimSpace(c.GetString("UserPwd", ""))
	newPwd := strings.TrimSpace(c.GetString("NewUserPwd", ""))
	confirmPwd := strings.TrimSpace(c.GetString("ConfirmPwd", ""))
	md5str := utils.String2md5(oldPwd)
	if oM.UserPwd != md5str {
		c.jsonResult(utils.JRCodeFailed, "原密码错误", "")
	}
	if len(newPwd) == 0 {
		c.jsonResult(utils.JRCodeFailed, "请输入新密码", "")
	}
	if newPwd != confirmPwd {
		c.jsonResult(utils.JRCodeFailed, "两次输入的新密码不一致", "")
	}
	oM.UserPwd = md5str*/
	o := orm.NewOrm()
	if _, err := o.Update(oM); err != nil {
		c.jsonResult(utils.JRCodeFailed, "保存失败", oM.Id)
	} else {
		c.jsonResult(utils.JRCodeSucc, "保存成功", oM.Id)
	}
}
