package models

import (
	"github.com/astaxie/beego/orm"
)

type Customer struct {
	Id     int
	UserId int
	Guid   string

	CustomerName string
	RealName     string

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

func (customer *Customer) TableName() string {
	return TableName("customer")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Customer))
}

func (customer *Customer) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(customer, fields...); err != nil {
		return err
	}
	return nil
}

func CustomerAdd(customer *Customer) (int64, error) {
	return orm.NewOrm().Insert(customer)
}

func CustomerGetById(id int) (*Customer, error) {
	u := new(Customer)

	err := orm.NewOrm().QueryTable(TableName("customer")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CustomerGetByName(customerName string) (*Customer, error) {
	u := new(Customer)

	err := orm.NewOrm().QueryTable(TableName("customer")).Filter("customer_name", customerName).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CustomerUpdate(customer *Customer, fields ...string) error {
	_, err := orm.NewOrm().Update(customer, fields...)
	return err
}

func CustomerDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("customer")).Filter("id", id).Delete()
	return err
}

func CustomerGetList(page, pageSize int) ([]*Customer, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Customer, 0)

	query := orm.NewOrm().QueryTable(TableName("customer"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
