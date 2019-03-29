package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ilus/models"
	"github.com/ilus/utils"
	"strconv"
	"strings"
)

//CategoryController 分类管理
type CategoryController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *CategoryController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "UpdateUrl")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

//Index 分类管理首页
func (c *CategoryController) Index() {
	c.Data["pageTitle"] = beego.AppConfig.String("appname") + " | 分类管理"
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = false
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "category/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "category/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("CategoryController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("CategoryController", "Delete")
}

// DataGrid 分类管理首页 表格获取数据
func (c *CategoryController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.CategoryQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.CategoryPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//Edit 添加、编辑标签界面
func (c *CategoryController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := models.Category{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.setTpl("category/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "category/edit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *CategoryController) Save() {
	var err error
	m := models.Category{}
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

}

//UpdateUrl 修改
func (c *CategoryController) UpdateUrl() {
	Id, _ := c.GetInt("pk", 0)
	oM, err := models.CategoryOne(Id)
	if err != nil || oM == nil {
		c.jsonResult(utils.JRCodeFailed, "选择的数据无效", 0)
	}
	value := c.GetString("value")
	oM.Url = value
	o := orm.NewOrm()
	if _, err := o.Update(oM); err == nil {
		c.jsonResult(utils.JRCodeSucc, "修改成功", oM.Id)
	} else {
		c.jsonResult(utils.JRCodeFailed, "修改失败", oM.Id)
	}
}

//Delete 批量删除
func (c *CategoryController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.CategoryDelete(ids); err == nil {
		c.jsonResult(utils.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(utils.JRCodeFailed, "删除失败", 0)
	}
}
