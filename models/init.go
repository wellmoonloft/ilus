package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(BackendUser), new(Resource), new(Role), new(RoleResourceRel), new(RoleBackendUserRel),
		new(Tags), new(Category), new(Article), new(ArticleCategoryTagsRel), new(BackendSettings), new(Comment))
}

// TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

// BackendUserTBName 获取 BackendUser 对应的表名称
func BackendUserTBName() string {
	return TableName("backend_user")
}

// ResourceTBName 获取 Resource 对应的表名称
func ResourceTBName() string {
	return TableName("resource")
}

// RoleTBName 获取 Role 对应的表名称
func RoleTBName() string {
	return TableName("role")
}

// RoleResourceRelTBName 角色与资源多对多关系表
func RoleResourceRelTBName() string {
	return TableName("role_resource_rel")
}

// RoleBackendUserRelTBName 角色与用户多对多关系表
func RoleBackendUserRelTBName() string {
	return TableName("role_backenduser_rel")
}

// TagsTBName 获取 Tags 对应的表名称
func TagsTBName() string {
	return TableName("tags")
}

// CategoryTBName 获取 Category 对应的表名称
func CategoryTBName() string {
	return TableName("category")
}

// CategoryTBName 获取 Category 对应的表名称
func ArticleTBName() string {
	return TableName("article")
}

// ArticleCategoryTagsRelTBName 角色与用户多对多关系表
func ArticleCategoryTagsRelTBName() string {
	return TableName("article_category_tags_rel")
}

// BackendSettingsTBName 常规系统设定表
func BackendSettingsTBName() string {
	return TableName("backend_settings")
}

// CommentTBName 常规系统设定表
func CommentTBName() string {
	return TableName("comment")
}
