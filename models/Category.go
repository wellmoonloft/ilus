package models

import "github.com/astaxie/beego/orm"

// TableName 设置表名
func (a *Category) TableName() string {
	return CategoryTBName()
}

// CategoryQueryParam 用于查询的类
type CategoryQueryParam struct {
	BaseQueryParam
	NameLike string //模糊查询

}

// Category 标签
type Category struct {
	Id   int
	Name string `orm:"size(64)"`
	Url  string `orm:"size(256)"`
	Abs  string `orm:"size(256)"`
}

// CategoryPageList 获取分页数据
func CategoryPageList(params *CategoryQueryParam) ([]*Category, int64) {
	query := orm.NewOrm().QueryTable(CategoryTBName())
	data := make([]*Category, 0)
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

// CategoryOne 根据id获取单条
func CategoryOne(id int) (*Category, error) {
	o := orm.NewOrm()
	m := Category{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// CategoryDelete 批量删除
func CategoryDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(CategoryTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
