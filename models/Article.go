package models

import "time"

// TableName 设置BackendUser表名
func (a *Article) TableName() string {
	return ArticleTBName()
}

// Article 标签
type Article struct {
	Id           int
	Name         string `orm:"size(256)"`
	Url          string `orm:"size(256)"`
	Value        string `orm:"type(text)"`
	IsComment    bool
	Pwd          string    `orm:"size(64)"`
	Date         time.Time `orm:"auto_now_add;type(datetime)"`
	Category     string    `orm:"size(64)"`
	Tags         string    `orm:"size(256)"`
	ThumbnailUrl string    `orm:"size(256)"`
}
