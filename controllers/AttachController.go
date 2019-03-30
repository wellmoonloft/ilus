package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ilus/imaging"
	"github.com/ilus/models"
	"github.com/ilus/utils"
	"io/ioutil"
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

	//读取附件库然后把结果塞进Data供前端展现
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

//UploadFile 上传文件
func (c *AttachController) UploadFile() {

	f, h, err := c.GetFile("fileImageUrl")
	if err != nil {
		c.jsonResult(utils.JRCodeFailed, "上传失败", "")
	}
	defer f.Close()
	filePath := "static/upload/" + h.Filename
	// 保存位置在 static/upload, 没有文件夹要先创建
	c.SaveToFile("fileImageUrl", filePath)

	//生成缩略图
	imgData, _ := ioutil.ReadFile(filePath)
	buf := bytes.NewBuffer(imgData)
	imagedecode, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	//生成缩略图，尺寸150*200，并保持到为文件2.jpg
	image := imaging.Resize(imagedecode, 256, 256, imaging.Lanczos)
	err = imaging.Save(image, "static/upload/"+"small-"+h.Filename)
	if err != nil {
		fmt.Println(err)
	}

	c.jsonResult(utils.JRCodeSucc, "上传成功", "/"+filePath)

}
