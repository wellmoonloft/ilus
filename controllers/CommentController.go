package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ilus/models"
	"github.com/ilus/utils"
	"strconv"
	"strings"
)

//CommentController 标签管理
type CommentController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *CommentController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "UpdateUrl")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

//Index 标签管理首页
func (c *CommentController) Index() {
	c.Data["pageTitle"] = beego.AppConfig.String("appname") + " | 评论管理"
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "comment/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "comment/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("CommentController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("CommentController", "Delete")
}

// DataGrid 标签管理首页 表格获取数据
func (c *CommentController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.CommentQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.CommentPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//Save 添加、编辑页面 保存
/*func (c *CommentController) Save() {
	var err error
	m := models.Comment{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(utils.JRCodeFailed, "获取数据失败", m.Id)
	}
	o := orm.NewOrm()
	if m.Id == 0 {
		if _, err = o.Insert(&m); err == nil {
			c.jsonResult(utils.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.jsonResult(utils.JRCodeFailed, "添加失败", m.Id)
		}

	} else {
		if _, err = o.Update(&m); err == nil {
			c.jsonResult(utils.JRCodeSucc, "编辑成功", m.Id)
		} else {
			c.jsonResult(utils.JRCodeFailed, "编辑失败", m.Id)
		}
	}

}*/

//Update 批量操作文件
func (c *CommentController) Update() {
	strs := c.GetString("ids")
	action := c.GetString("action")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.CommentUpate(ids, action); err == nil {
		c.jsonResult(utils.JRCodeSucc, fmt.Sprintf("成功修改 %d 项", num), 0)
	} else {
		c.jsonResult(utils.JRCodeFailed, "修改失败", 0)
	}
}

//Delete 批量删除
func (c *CommentController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.CommentDelete(ids); err == nil {
		c.jsonResult(utils.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(utils.JRCodeFailed, "删除失败", 0)
	}
}
