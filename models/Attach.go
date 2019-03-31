package models

import (
	"github.com/astaxie/beego/orm"
	"math"
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
	Id        int
	Name      string    `orm:"size(64)"`
	Type      string    `orm:"size(64)"`
	Url       string    `orm:"size(256)"`
	Thumbnail string    `orm:"size(256)"`
	Size      string    `orm:"size(64)"`
	ImgSize   string    `orm:"size(64)"`
	Date      time.Time `orm:"auto_now_add;type(datetime)"`
}

// AttachPageList 获取分页数据
func AttachPageList(params *AttachQueryParam) ([]*Attach, map[string]interface{}) {
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

	res := Paginator(1, 12, int64(total))
	params.Limit = 12

	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	//return data, total
	return data, res
}

//Paginator 分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, prepage int, nums int64) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	return paginatorMap
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
