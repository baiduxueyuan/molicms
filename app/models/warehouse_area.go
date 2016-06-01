package models

import (
	"github.com/astaxie/beego/orm"
)

type WarehouseArea struct {
	Id                int
	UserId            int
	Guid              string
	WarehouseId       int
	WarehouseAreaName string

	CreateTime int64
	ModifyTime int64

	Remark string
	Status int
}

func (warehousearea *WarehouseArea) TableName() string {
	return TableName("warehouse_area")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(WarehouseArea))
}

func (warehousearea *WarehouseArea) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(warehousearea, fields...); err != nil {
		return err
	}
	return nil
}

func WarehouseAreaAdd(warehousearea *WarehouseArea) (int64, error) {
	return orm.NewOrm().Insert(warehousearea)
}

func WarehouseAreaGetById(id int) (*WarehouseArea, error) {
	u := new(WarehouseArea)

	err := orm.NewOrm().QueryTable(TableName("warehouse_area")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func WarehouseAreaGetByName(warehouseareaName string) (*WarehouseArea, error) {
	u := new(WarehouseArea)

	err := orm.NewOrm().QueryTable(TableName("warehouse_area")).Filter("warehouse_area_name", warehouseareaName).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func WarehouseAreaUpdate(warehousearea *WarehouseArea, fields ...string) error {
	_, err := orm.NewOrm().Update(warehousearea, fields...)
	return err
}

func WarehouseAreaDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("warehouse_area")).Filter("id", id).Delete()
	return err
}

func WarehouseAreaGetList(page, pageSize int) ([]*WarehouseArea, int64) {
	offset := (page - 1) * pageSize

	list := make([]*WarehouseArea, 0)

	query := orm.NewOrm().QueryTable(TableName("warehouse_area"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
