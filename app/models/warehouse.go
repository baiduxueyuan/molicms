package models

import (
	"github.com/astaxie/beego/orm"
)

type Warehouse struct {
	Id     int
	UserId int
	Guid   string

	WarehouseName string
	CountryId     int
	ProvinceId    int
	CityId        int
	AreaId        int
	Country       string
	Province      string
	City          string
	Area          string
	Address       string

	Phone   string
	QQ      string
	Weixin  string
	Mobile  string
	Manager string

	CreateTime int64
	ModifyTime int64

	Remark string
	Status int
}

func (warehouse *Warehouse) TableName() string {
	return TableName("warehouse")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Warehouse))
}

func (warehouse *Warehouse) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(warehouse, fields...); err != nil {
		return err
	}
	return nil
}

func WarehouseAdd(warehouse *Warehouse) (int64, error) {
	return orm.NewOrm().Insert(warehouse)
}

func WarehouseGetById(id int) (*Warehouse, error) {
	u := new(Warehouse)

	err := orm.NewOrm().QueryTable(TableName("warehouse")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func WarehouseGetByName(warehouseName string) (*Warehouse, error) {
	u := new(Warehouse)

	err := orm.NewOrm().QueryTable(TableName("warehouse")).Filter("warehouse_name", warehouseName).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func WarehouseUpdate(warehouse *Warehouse, fields ...string) error {
	_, err := orm.NewOrm().Update(warehouse, fields...)
	return err
}

func WarehouseDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("warehouse")).Filter("id", id).Delete()
	return err
}

func WarehouseGetList(page, pageSize int) ([]*Warehouse, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Warehouse, 0)

	query := orm.NewOrm().QueryTable(TableName("warehouse"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
