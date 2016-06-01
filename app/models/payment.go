package models

import (
	"github.com/astaxie/beego/orm"
)

type Payment struct {
	Id     int
	UserId int
	Guid   string

	PaymentName string
	RealName    string

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

func (payment *Payment) TableName() string {
	return TableName("payment")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Payment))
}

func (payment *Payment) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(payment, fields...); err != nil {
		return err
	}
	return nil
}

func PaymentAdd(payment *Payment) (int64, error) {
	return orm.NewOrm().Insert(payment)
}

func PaymentGetById(id int) (*Payment, error) {
	u := new(Payment)

	err := orm.NewOrm().QueryTable(TableName("payment")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func PaymentGetByName(paymentName string) (*Payment, error) {
	u := new(Payment)

	err := orm.NewOrm().QueryTable(TableName("payment")).Filter("payment_name", paymentName).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func PaymentUpdate(payment *Payment, fields ...string) error {
	_, err := orm.NewOrm().Update(payment, fields...)
	return err
}

func PaymentDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("payment")).Filter("id", id).Delete()
	return err
}

func PaymentGetList(page, pageSize int) ([]*Payment, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Payment, 0)

	query := orm.NewOrm().QueryTable(TableName("payment"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
