package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Order struct {
	Id         int
	Guid       string
	StroeId    int
	UserId     int
	CustomerId int
	//订单基本信息
	OrderName string
	//
	OrderCustomId string
	OrderContent  string
	OrderCount    int
	OrderAmount   float64
	//订单参数
	//地区
	Province   string
	ProvinceId int
	City       string
	CityId     int
	Area       string
	AreaId     int
	Status     int
	//时间
	CreateTime int64
	ModifyTime int64
}

func (t *Order) TableName() string {
	return TableName("order")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Order))
}

func (t *Order) Update(fields ...string) error {
	if t.OrderName == "" {
		return fmt.Errorf("订单名称不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func OrderAdd(obj *Order) (int64, error) {
	if obj.OrderName == "" {
		return 0, fmt.Errorf("订单名称不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func OrderGetById(id int) (*Order, error) {
	obj := &Order{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func OrderDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("order")).Filter("id", id).Delete()
	return err
}

func OrderGetList(page, pageSize int) ([]*Order, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Order, 0)

	query := orm.NewOrm().QueryTable(TableName("order"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
