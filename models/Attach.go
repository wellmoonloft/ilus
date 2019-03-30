package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// TableName 设置表名
func (a *Attach) TableName() string {
	return AttachTBName()
}

// AttachQueryParam 用于查询的类
type AttachQueryParam struct {
	BaseQueryParam
	NameLike string //查询
}

// Attach 标签
type Attach struct {
	Id      int
	Name    string    `orm:"size(64)"`
	Type    string    `orm:"size(64)"`
	Url     string    `orm:"size(256)"`
	Size    string    `orm:"size(64)"`
	ImgSize string    `orm:"size(64)"`
	Date    time.Time `orm:"auto_now_add;type(datetime)"`
}

// AttachPageList 获取分页数据
func AttachPageList(params *AttachQueryParam) ([]*Attach, int64) {
	query := orm.NewOrm().QueryTable(AttachTBName())
	data := make([]*Attach, 0)
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

// AttachOne 根据id获取单条
func AttachOne(id int) (*Attach, error) {
	o := orm.NewOrm()
	m := Attach{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// AttachDelete 批量删除
func AttachDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(AttachTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
