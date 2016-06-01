package models

import (
	"github.com/astaxie/beego/orm"
)

type Refund struct {
	Id     int
	UserId int
	Guid   string

	RefundName string
	RealName   string

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

func (refund *Refund) TableName() string {
	return TableName("refund")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Refund))
}

func (refund *Refund) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(refund, fields...); err != nil {
		return err
	}
	return nil
}

func RefundAdd(refund *Refund) (int64, error) {
	return orm.NewOrm().Insert(refund)
}

func RefundGetById(id int) (*Refund, error) {
	u := new(Refund)

	err := orm.NewOrm().QueryTable(TableName("refund")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func RefundGetByName(refundName string) (*Refund, error) {
	u := new(Refund)

	err := orm.NewOrm().QueryTable(TableName("refund")).Filter("refund_name", refundName).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func RefundUpdate(refund *Refund, fields ...string) error {
	_, err := orm.NewOrm().Update(refund, fields...)
	return err
}

func RefundDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("refund")).Filter("id", id).Delete()
	return err
}

func RefundGetList(page, pageSize int) ([]*Refund, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Refund, 0)

	query := orm.NewOrm().QueryTable(TableName("refund"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
