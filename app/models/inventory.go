package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Inventory struct {
	Id      int
	StroeId int
	UserId  int
	Guid    string
	//SEO
	Keywords    string
	Description string
	//仓库信息
	WarehouseId       int    //仓库ID
	WarehouseName     string //仓库名称
	WarehouseAreaName string //仓库库区
	WarehouseAreaId   int    //仓库库区ID
	StockLocationId   int    //库位ID
	StockLocation     string //库位  库、架、层、位
	//产品基本信息
	ProductName  string //产品名称
	SerialNumber string //序列号
	Barcode      string //条码
	RFID         string //RFID
	CustomID     string //自定义ID
	//产品参数
	Type   string //产品型号
	Color  string //产品颜色
	Size   string //产品尺寸
	Weight int    //产品重量
	Unit   string //产品单位

	//库存
	Inventory      int //总库存
	InventoryLock  int //锁定库存
	InventoryAlarm int //告警库存
	//统计
	SaleCount int //销售量
	//other
	Level int //搜索排序等级
	//时间
	CreateTime int64
	ModifyTime int64
}

func (t *Inventory) TableName() string {
	return TableName("inventory")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Inventory))
}

func (t *Inventory) Update(fields ...string) error {
	if t.ProductName == "" {
		return fmt.Errorf("产品名称不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func InventoryAdd(obj *Inventory) (int64, error) {
	if obj.ProductName == "" {
		return 0, fmt.Errorf("产品名称不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func InventoryGetById(id int) (*Inventory, error) {
	obj := &Inventory{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func InventoryDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("inventory")).Filter("id", id).Delete()
	return err
}

func InventoryGetList(page, pageSize int) ([]*Inventory, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Inventory, 0)

	query := orm.NewOrm().QueryTable(TableName("inventory"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
