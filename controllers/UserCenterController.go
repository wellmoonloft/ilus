package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ilus/models"
	"github.com/ilus/utils"
	"strings"
)

//UserCenterController 后台用户管理
type UserCenterController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *UserCenterController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	c.checkLogin()
}

//Profile 后台用户管理首页
func (c *UserCenterController) Profile() {
	c.Data["pageTitle"] = beego.AppConfig.String("appname") + " | 个人信息"
	Id := c.curUser.Id
	m, err := models.BackendUserOne(Id)
	if m == nil || err != nil {
		c.pageError("数据无效，请刷新后重试")
	}
	c.Data["hasAvatar"] = len(m.Avatar) > 0
	utils.LogDebug(m.Avatar)
	c.Data["m"] = m
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "usercenter/profile_headcssjs.html"
	c.LayoutSections["footerjs"] = "usercenter/profile_footerjs.html"
}

//BasicInfoSave 基础信息保存
func (c *UserCenterController) BasicInfoSave() {
	Id := c.curUser.Id
	oM, err := models.BackendUserOne(Id)
	if oM == nil || err != nil {
		c.jsonResult(utils.JRCodeFailed, "数据无效，请刷新后重试", "")
	}
	m := models.BackendUser{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(utils.JRCodeFailed, "获取数据失败", m.Id)
	}
	oM.RealName = m.RealName
	oM.Mobile = m.Mobile
	oM.Email = m.Email
	oM.Avatar = c.GetString("ImageUrl")
	o := orm.NewOrm()
	if _, err := o.Update(oM); err != nil {
		c.jsonResult(utils.JRCodeFailed, "编辑失败", m.Id)
	} else {
		c.setBackendUser2Session(Id)
		c.jsonResult(utils.JRCodeSucc, "保存成功", m.Id)
	}
}

//PasswordSave 密码信息保存
func (c *UserCenterController) PasswordSave() {
	Id := c.curUser.Id
	oM, err := models.BackendUserOne(Id)
	if oM == nil || err != nil {
		c.pageError("数据无效，请刷新后重试")
	}
	oldPwd := strings.TrimSpace(c.GetString("UserPwd", ""))
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
	oM.UserPwd = md5str
	o := orm.NewOrm()
	if _, err := o.Update(oM); err != nil {
		c.jsonResult(utils.JRCodeFailed, "保存失败", oM.Id)
	} else {
		c.setBackendUser2Session(Id)
		c.jsonResult(utils.JRCodeSucc, "保存成功", oM.Id)
	}
}

//UploadImage 头像信息保存
func (c *UserCenterController) UploadImage() {
	//这里type没有用，只是为了演示传值
	stype, _ := c.GetInt32("type", 0)
	if stype > 0 {
		f, h, err := c.GetFile("fileImageUrl")
		if err != nil {
			c.jsonResult(utils.JRCodeFailed, "上传失败", "")
		}
		defer f.Close()
		filePath := "static/upload/" + h.Filename
		// 保存位置在 static/upload, 没有文件夹要先创建
		c.SaveToFile("fileImageUrl", filePath)
		c.jsonResult(utils.JRCodeSucc, "上传成功", "/"+filePath)
	} else {
		c.jsonResult(utils.JRCodeFailed, "上传失败", "")
	}
}
