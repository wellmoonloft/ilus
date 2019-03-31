package routers

import (
	"github.com/astaxie/beego"
	"github.com/ilus/controllers"
)

func init() {

	//用户角色路由
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/role/datagrid", &controllers.RoleController{}, "Get,Post:DataGrid")
	beego.Router("/role/edit/?:id", &controllers.RoleController{}, "Get,Post:Edit")
	beego.Router("/role/delete", &controllers.RoleController{}, "Post:Delete")
	beego.Router("/role/datalist", &controllers.RoleController{}, "Post:DataList")
	beego.Router("/role/allocate", &controllers.RoleController{}, "Post:Allocate")
	beego.Router("/role/updateseq", &controllers.RoleController{}, "Post:UpdateSeq")

	//标签路由
	beego.Router("/tags/index", &controllers.TagsController{}, "*:Index")
	beego.Router("/tags/datagrid", &controllers.TagsController{}, "Get,Post:DataGrid")
	beego.Router("/tags/edit/?:id", &controllers.TagsController{}, "Get,Post:Edit")
	beego.Router("/tags/delete", &controllers.TagsController{}, "Post:Delete")
	beego.Router("/tags/updateUrl", &controllers.TagsController{}, "Post:UpdateUrl")

	//分类路由
	beego.Router("/category/index", &controllers.CategoryController{}, "*:Index")
	beego.Router("/category/datagrid", &controllers.CategoryController{}, "Get,Post:DataGrid")
	beego.Router("/category/edit/?:id", &controllers.CategoryController{}, "Get,Post:Edit")
	beego.Router("/category/delete", &controllers.CategoryController{}, "Post:Delete")
	beego.Router("/category/updateUrl", &controllers.CategoryController{}, "Post:UpdateUrl")

	//评论路由
	beego.Router("/comment/index", &controllers.CommentController{}, "*:Index")
	beego.Router("/comment/datagrid", &controllers.CommentController{}, "Get,Post:DataGrid")
	beego.Router("/comment/delete", &controllers.CommentController{}, "Post:Delete")
	beego.Router("/comment/update", &controllers.CommentController{}, "Post:Update")

	//资源路由
	beego.Router("/resource/index", &controllers.ResourceController{}, "*:Index")
	beego.Router("/resource/treegrid", &controllers.ResourceController{}, "POST:TreeGrid")
	beego.Router("/resource/edit/?:id", &controllers.ResourceController{}, "Get,Post:Edit")
	beego.Router("/resource/parent", &controllers.ResourceController{}, "Post:ParentTreeGrid")
	beego.Router("/resource/delete", &controllers.ResourceController{}, "Post:Delete")
	//快速修改顺序
	beego.Router("/resource/updateseq", &controllers.ResourceController{}, "Post:UpdateSeq")
	//通用选择面板
	beego.Router("/resource/select", &controllers.ResourceController{}, "Get:Select")
	//用户有权管理的菜单列表（包括区域）
	beego.Router("/resource/usermenutree", &controllers.ResourceController{}, "POST:UserMenuTree")
	beego.Router("/resource/checkurlfor", &controllers.ResourceController{}, "POST:CheckUrlFor")

	//后台用户路由
	beego.Router("/backenduser/index", &controllers.BackendUserController{}, "*:Index")
	beego.Router("/backenduser/datagrid", &controllers.BackendUserController{}, "POST:DataGrid")
	beego.Router("/backenduser/edit/?:id", &controllers.BackendUserController{}, "Get,Post:Edit")
	beego.Router("/backenduser/delete", &controllers.BackendUserController{}, "Post:Delete")

	//后台用户中心
	beego.Router("/usercenter/profile", &controllers.UserCenterController{}, "Get:Profile")
	beego.Router("/usercenter/basicinfosave", &controllers.UserCenterController{}, "Post:BasicInfoSave")
	beego.Router("/usercenter/uploadimage", &controllers.UserCenterController{}, "Post:UploadImage")
	beego.Router("/usercenter/passwordsave", &controllers.UserCenterController{}, "Post:PasswordSave")

	//文章管理
	beego.Router("/article/new", &controllers.ArticleController{}, "*:New")
	beego.Router("/article/index", &controllers.ArticleController{}, "*:Index")
	beego.Router("/article/datagrid", &controllers.ArticleController{}, "Get,Post:DataGrid")
	beego.Router("/article/edit/?:id", &controllers.ArticleController{}, "Get,Post:Edit")
	beego.Router("/article/delete", &controllers.ArticleController{}, "Post:Delete")
	beego.Router("/article/updateUrl", &controllers.ArticleController{}, "Post:UpdateUrl")

	//附件路由
	beego.Router("/attach/index", &controllers.AttachController{}, "*:Index")
	beego.Router("/attach/uploadfile", &controllers.AttachController{}, "Post:UploadFile")
	beego.Router("/attach/refresh/?:action", &controllers.AttachController{}, "Get:ReFresh")

	//系统设定
	beego.Router("/backendsettings/index", &controllers.BackendSettingsController{}, "*:Index")
	beego.Router("/backendsettings/generalSave", &controllers.BackendSettingsController{}, "Post:GeneralSave")

	//首页
	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/login", &controllers.HomeController{}, "*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")
	beego.Router("/home/logout", &controllers.HomeController{}, "*:Logout")

	beego.Router("/home/404", &controllers.HomeController{}, "*:Page404")
	beego.Router("/home/error/?:error", &controllers.HomeController{}, "*:Error")

	beego.Router("/", &controllers.HomeController{}, "*:Index")

}
