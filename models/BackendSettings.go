package models

import "github.com/astaxie/beego/orm"

// TableName 设置表名
func (a *BackendSettings) TableName() string {
	return BackendSettingsTBName()
}

// BackendSettingsQueryParam 用于查询的类
type BackendSettingsQueryParam struct {
	BaseQueryParam
	NameLike string //模糊查询
}

// BackendSettings 标签
type BackendSettings struct {
	Id int
	//general
	Name     string `orm:"size(64)"`
	Url      string `orm:"size(128)"`
	Language int
	Logo     string `orm:"size(128)"`
	Favicon  string `orm:"size(128)"`
	Footer   string `orm:"size(256)"`
	//seo
	Key           string `orm:"size(64)"`
	Abs           string `orm:"size(128)"`
	BaiduToken    string `orm:"size(128)"`
	GoogleCaptcha string `orm:"size(128)"`
	BingCaptcha   string `orm:"size(128)"`
	BaiduCaptcha  string `orm:"size(128)"`
	//cagegory
	//comment
	//attach
	AttachPath   int
	AttachDomain string `orm:"size(128)"`
	AccessKey    string `orm:"size(128)"`
	SecretKey    string `orm:"size(128)"`
	Bucket       string `orm:"size(128)"`
	Strategy     string `orm:"size(64)"`
	//admin
	//email
	IsEnable  bool
	Smtp      string `orm:"size(128)"`
	EmailUser string `orm:"size(64)"`
	EmailPwd  string `orm:"size(64)"`
	SendName  string `orm:"size(64)"`
	//other
	ApiIsEnable bool
	ApiToken    string `orm:"size(256)"`
	StatsCode   string `orm:"size(256)"`
}

// BackendSettingsPageList 获取分页数据
func BackendSettingsPageList(params *BackendSettingsQueryParam) ([]*BackendSettings, int64) {
	query := orm.NewOrm().QueryTable(BackendSettingsTBName())
	data := make([]*BackendSettings, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	/*query = query.Filter("username__istartswith", params.NameLike)
	query = query.Filter("realname__istartswith", params.RealNameLike)
	if len(params.Mobile) > 0 {
		query = query.Filter("mobile", params.Mobile)
	}
	if len(params.SearchStatus) > 0 {
		query = query.Filter("status", params.SearchStatus)
	}*/
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// BackendSettingsOne 根据id获取单条
func BackendSettingsOne(id int) (*BackendSettings, error) {
	o := orm.NewOrm()
	m := BackendSettings{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// BackendSettingsOneByName 根据用户名密码获取单条
func BackendSettingsOneByName(name string) (*BackendSettings, error) {
	m := BackendSettings{}
	err := orm.NewOrm().QueryTable(BackendSettingsTBName()).Filter("name", name).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
