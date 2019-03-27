package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// TableName 设置BackendUser表名
func (a *Article) TableName() string {
	return ArticleTBName()
}

// ArticleQueryParam 用于查询的类
type ArticleQueryParam struct {
	BaseQueryParam
	NameLike string //模糊查询

}

// Article 标签
type Article struct {
	Id           int
	Name         string `orm:"size(256)"`
	Url          string `orm:"size(256)"`
	Value        string `orm:"type(text)"`
	IsComment    bool
	IsPush       bool
	Visits       int32
	Comment      int32
	Pwd          string    `orm:"size(64)"`
	Date         time.Time `orm:"auto_now_add;type(datetime)"`
	Category     string    `orm:"size(64)"`
	Tags         string    `orm:"size(256)"`
	ThumbnailUrl string    `orm:"size(256)"`
}

// ArticlePageList 获取分页数据
func ArticlePageList(params *ArticleQueryParam) ([]*Article, int64) {
	query := orm.NewOrm().QueryTable(ArticleTBName())
	data := make([]*Article, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("name__istartswith", params.NameLike)

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// ArticleOne 根据id获取单条
func ArticleOne(id int) (*Article, error) {
	o := orm.NewOrm()
	m := Article{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// ArticleDelete 批量删除
func ArticleDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(ArticleTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
