package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/ilus/models"
	"github.com/ilus/utils"
)

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
func (c *BackendSettingsController) UploadImage() {
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
