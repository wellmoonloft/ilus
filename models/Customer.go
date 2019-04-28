package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// TableName 设置BackendUser表名
func (a *Customer) TableName() string {
	return CustomerTBName()
}

// CustomerQueryParam 用于查询的类
type CustomerQueryParam struct {
	BaseQueryParam
	RealNameLike string //模糊查询
	CardCodeLike string //模糊查询
	Mobile       string //精确查询
	SearchStatus string //为空不查询，有值精确查询
}

// Customer 实体类
type Customer struct {
	Id               int
	Code             string
	IDCard           string
	RealName         string `orm:"size(32)"`
	WeChatName       string `orm:"size(32)"`
	NickName         string `orm:"size(32)"`
	CardCode         string
	CardLevel        int8
	Pwd              string
	Sexy             bool
	Status           int8
	Purchase         int32
	PurchaseTimes    int32
	LastPurchaseDate time.Time
	Credit           int32
	Source           string
	Mobile           string `orm:"size(16)"`
	Email            string `orm:"size(256)"`
	Avatar           string `orm:"size(256)"`
	JoinDate         time.Time
	LoseDate         time.Time
}

// CustomerPageList 获取分页数据
func CustomerPageList(params *CustomerQueryParam) ([]*Customer, int64) {
	query := orm.NewOrm().QueryTable(CustomerTBName())
	data := make([]*Customer, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("realname__istartswith", params.RealNameLike)
	query = query.Filter("cardcode__istartswith", params.CardCodeLike)
	if len(params.Mobile) > 0 {
		query = query.Filter("mobile", params.Mobile)
	}
	if len(params.SearchStatus) > 0 {
		query = query.Filter("status", params.SearchStatus)
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// CustomerOne 根据id获取单条
func CustomerOne(id int) (*Customer, error) {
	o := orm.NewOrm()
	m := Customer{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// CustomerOneByUserName 根据用户名密码获取单条
func CustomerOneByUserName(username, userpwd string) (*Customer, error) {
	m := Customer{}
	err := orm.NewOrm().QueryTable(CustomerTBName()).Filter("username", username).Filter("userpwd", userpwd).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
