package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// TableName 设置表名
func (a *Comment) TableName() string {
	return CommentTBName()
}

// CommentQueryParam 用于查询的类
type CommentQueryParam struct {
	BaseQueryParam
	UserLike string //模糊查询
}

// Comment 标签
type Comment struct {
	Id      int
	User    string    `orm:"size(64)"`
	UserUrl string    `orm:"size(256)"`
	Value   string    `orm:"size(256)"`
	Article string    `orm:"size(256)"`
	Date    time.Time `orm:"auto_now_add;type(datetime)"`
}

// CommentPageList 获取分页数据
func CommentPageList(params *CommentQueryParam) ([]*Comment, int64) {
	query := orm.NewOrm().QueryTable(CommentTBName())
	data := make([]*Comment, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("user__istartswith", params.UserLike)

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// CommentOne 根据id获取单条
func CommentOne(id int) (*Comment, error) {
	o := orm.NewOrm()
	m := Comment{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// CommentDelete 批量删除
func CommentDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(CommentTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
