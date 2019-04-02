package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ilus/imaging"
	"github.com/ilus/models"
	"github.com/ilus/utils"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

//AttachController 附件管理
type AttachController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *AttachController) Prepare() {
	c.BaseController.Prepare()
	c.checkLogin()
}

//Index 附件管理首页
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
	//数据
	result["rows"] = data
	c.Data["json"] = result

	//c.ServeJSON()

}

//ReFresh 附件管理首页
func (c *AttachController) ReFresh() {

	action := c.GetString(":action")
	//读取附件库然后把结果塞进Data供前端展现
	var params models.AttachQueryParam
	/*switch action {
	case "first":
		...
	case "last":

	case "next":
		...
	default:
		...
	}*/

	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.AttachPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	//总页数
	result["pages"] = math.Ceil(float64(total) / float64(12))
	//数据
	result["rows"] = data
	//当前页数
	result["currpage"] = action
	//首页是否可操作
	result["isfirst"] = true
	//上页是否可操作
	result["islast"] = true
	//下页是否可操作
	result["isnext"] = false
	//尾页是否可操作
	result["isend"] = false

	c.Data["json"] = result

	//c.ServeJSON()

}

//UploadFile 上传文件
func (c *AttachController) UploadFile() {

	f, h, err := c.GetFile("fileImageUrl")
	if err != nil {
		c.jsonResult(utils.JRCodeFailed, "上传失败", "")
	}
	defer f.Close()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	filetype := h.Header.Get("Content-Type")

	//判断文件夹是否存在并自动创建文件夹
	utils.PathExists("static/upload/" + year + "/")
	utils.PathExists("static/upload/" + year + "/" + month + "/")
	utils.PathExists("static/upload/" + year + "/" + month + "/file/")
	utils.PathExists("static/upload/" + year + "/" + month + "/thumbnail/")
	//部署的时候记得加判断条件，如果建文件夹权限不足

	//把文件名改为当前时间戳
	filePath := "static/upload/" + year + "/" + month + "/file/" + timestamp + h.Filename
	c.SaveToFile("fileImageUrl", filePath)
	//读取图片文件并转码
	imgData, _ := ioutil.ReadFile(filePath)
	buf := bytes.NewBuffer(imgData)
	imagedecode, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	//生成缩略图，尺寸256*256，并保存
	thumbnail := "static/upload/" + year + "/" + month + "/thumbnail/small-" + timestamp + h.Filename
	image := imaging.Resize(imagedecode, 256, 256, imaging.Lanczos)
	err = imaging.Save(image, thumbnail)
	if err != nil {
		fmt.Println(err)
	}

	//保存到数据库
	m := models.Attach{}
	o := orm.NewOrm()
	m.Name = timestamp + h.Filename
	m.Date = time.Now()
	m.ImgSize = strconv.FormatInt(h.Size, 10)
	m.Url = filePath
	m.Thumbnail = thumbnail
	m.Type = filetype

	if _, err := o.Insert(&m); err == nil {
		c.jsonResult(utils.JRCodeSucc, "添加成功", m.Id)
	} else {
		c.jsonResult(utils.JRCodeFailed, "添加失败", m.Id)
	}

	c.jsonResult(utils.JRCodeSucc, "上传成功", "/"+filePath)

}

//Save 添加、编辑页面 保存
func (c *AttachController) Save() {

	m := models.Attach{}
	o := orm.NewOrm()
	if _, err := o.Insert(&m); err == nil {
		c.jsonResult(utils.JRCodeSucc, "添加成功", m.Id)
	} else {
		c.jsonResult(utils.JRCodeFailed, "添加失败", m.Id)
	}

}

//Delete 批量删除
func (c *AttachController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.AttachDelete(ids); err == nil {
		c.jsonResult(utils.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(utils.JRCodeFailed, "删除失败", 0)
	}
}
