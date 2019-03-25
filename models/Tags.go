package models

import (
	"github.com/astaxie/beego/orm"
)

// TableName 设置表名
func (a *Tags) TableName() string {
	return TagsTBName()
}

// TagsQueryParam 用于查询的类
type TagsQueryParam struct {
	BaseQueryParam
	NameLike string //模糊查询

}

// Tags 标签
type Tags struct {
	Id   int
	Name string `orm:"size(64)"`
	Url  string `orm:"size(256)"`
}

// TagsPageList 获取分页数据
func TagsPageList(params *TagsQueryParam) ([]*Tags, int64) {
	query := orm.NewOrm().QueryTable(TagsTBName())
	data := make([]*Tags, 0)
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

// TagsOne 根据id获取单条
func TagsOne(id int) (*Tags, error) {
	o := orm.NewOrm()
	m := Tags{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// TagsDelete 批量删除
func TagsDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(TagsTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
