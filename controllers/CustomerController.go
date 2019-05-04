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

//CustomerController 后台用户管理
type CustomerController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *CustomerController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "GetCustomerOne")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()

}

//Index 用户管理首页
func (c *CustomerController) Index() {
	c.Data["pageTitle"] = beego.AppConfig.String("appname") + " | 会员管理"
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "customer/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "customer/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("CustomerController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("CustomerController", "Delete")
}

// DataGrid 标签管理首页 表格获取数据
func (c *CustomerController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.CustomerQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.CustomerPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *CustomerController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := &models.Customer{}
	var err error
	if Id > 0 {
		m, err = models.CustomerOne(Id)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	//获取关联的roleId列表
	/*var roleIds []string
	for _, item := range m.RoleBackendUserRel {
		roleIds = append(roleIds, strconv.Itoa(item.Role.Id))
	}
	c.Data["roles"] = strings.Join(roleIds, ",")*/
	c.setTpl("customer/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "customer/edit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *CustomerController) Save() {
	m := models.Customer{}
	o := orm.NewOrm()
	var err error
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(utils.JRCodeFailed, "获取数据失败", m.Id)
	}
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

//Delete 批量删除
func (c *CustomerController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	query := orm.NewOrm().QueryTable(models.CustomerTBName())
	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		c.jsonResult(utils.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(utils.JRCodeFailed, "删除失败", 0)
	}
}

// GetCustomerOne 添加 编辑 页面
func (c *CustomerController) GetCustomerOne() {

	Id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	m := &models.Customer{}
	var err error
	if Id > 0 {
		m, err = models.CustomerOne(Id)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["json"] = m
	c.ServeJSON()
}
