package models

import (
	"github.com/astaxie/beego/orm"
)

type Aftersale struct {
	Id     int
	UserId int
	Guid   string

	AftersaleName string
	RealName      string

	tags string

	CountryId  int
	ProvinceId int
	CityId     int
	AreaId     int
	Country    string
	Province   string
	City       string
	Area       string
	Address    string

	Phone   string
	Mobile  string
	Mobile2 string
	Email   string
	Weixin  string
	QQ      string
	QQ2     string

	BuyCount  int
	BuyAmount float64

	FirstFrom          string
	FirstFromStroeId   int
	FirstFromProductId int

	MobileType string
	PCType     string

	BuyProduct string

	CreateTime int64
	ModifyTime int64

	Remark string
	Status int
}

func (aftersale *Aftersale) TableName() string {
	return TableName("aftersale")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Aftersale))
}

func (aftersale *Aftersale) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(aftersale, fields...); err != nil {
		return err
	}
	return nil
}

func AftersaleAdd(aftersale *Aftersale) (int64, error) {
	return orm.NewOrm().Insert(aftersale)
}

func AftersaleGetById(id int) (*Aftersale, error) {
	u := new(Aftersale)

	err := orm.NewOrm().QueryTable(TableName("aftersale")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func AftersaleGetByName(aftersaleName string) (*Aftersale, error) {
	u := new(Aftersale)

	err := orm.NewOrm().QueryTable(TableName("aftersale")).Filter("aftersale_name", aftersaleName).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func AftersaleUpdate(aftersale *Aftersale, fields ...string) error {
	_, err := orm.NewOrm().Update(aftersale, fields...)
	return err
}

func AftersaleDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("aftersale")).Filter("id", id).Delete()
	return err
}

func AftersaleGetList(page, pageSize int) ([]*Aftersale, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Aftersale, 0)

	query := orm.NewOrm().QueryTable(TableName("aftersale"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
