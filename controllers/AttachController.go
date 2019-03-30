package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/ilus/models"
)

//AttachController 标签管理
type AttachController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *AttachController) Prepare() {
	c.BaseController.Prepare()
	c.checkLogin()
}

//Index 标签管理首页
func (c *AttachController) Index() {
	c.Data["pageTitle"] = beego.AppConfig.String("appname") + " | 附件管理"
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "attach/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "attach/index_footerjs.html"
	//页面里按钮权限控制
	/*c.Data["canEdit"] = c.checkActionAuthor("CommentController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("CommentController", "Delete")*/

	var params models.AttachQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.AttachPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result

}

// DataGrid 标签管理首页 表格获取数据
func (c *AttachController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.AttachQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.AttachPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
