package models

import "time"

// TableName 设置表名
func (a *ArticleCategoryTagsRel) TableName() string {
	return ArticleCategoryTagsRelTBName()
}

// ArticleCategoryTagsRel 文章、分类、标签关系
type ArticleCategoryTagsRel struct {
	Id       int
	Article  *Article  `orm:"rel(fk)"`  //外键
	Category *Category `orm:"rel(fk)" ` // 外键
	Tags     *Tags     `orm:"rel(fk)" ` // 外键
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}
