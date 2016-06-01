package models

import (
	"github.com/astaxie/beego/orm"
)

type WarehouseLocation struct {
	Id                    int
	UserId                int
	Guid                  string
	WarehouseId           int
	WarehouseAreaId       int
	WarehouseLocationName string

	CreateTime int64
	ModifyTime int64

	Remark string
	Status int
}

func (warehouselocation *WarehouseLocation) TableName() string {
	return TableName("warehouse_location")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(WarehouseLocation))
}

func (warehouselocation *WarehouseLocation) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(warehouselocation, fields...); err != nil {
		return err
	}
	return nil
}

func WarehouseLocationAdd(warehouselocation *WarehouseLocation) (int64, error) {
	return orm.NewOrm().Insert(warehouselocation)
}

func WarehouseLocationGetById(id int) (*WarehouseLocation, error) {
	u := new(WarehouseLocation)

	err := orm.NewOrm().QueryTable(TableName("warehouse_location")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func WarehouseLocationGetByName(warehouselocationName string) (*WarehouseLocation, error) {
	u := new(WarehouseLocation)

	err := orm.NewOrm().QueryTable(TableName("warehouse_location")).Filter("warehouse_location_name", warehouselocationName).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func WarehouseLocationUpdate(warehouselocation *WarehouseLocation, fields ...string) error {
	_, err := orm.NewOrm().Update(warehouselocation, fields...)
	return err
}

func WarehouseLocationDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("warehouse_location")).Filter("id", id).Delete()
	return err
}

func WarehouseLocationGetList(page, pageSize int) ([]*WarehouseLocation, int64) {
	offset := (page - 1) * pageSize

	list := make([]*WarehouseLocation, 0)

	query := orm.NewOrm().QueryTable(TableName("warehouse_location"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
