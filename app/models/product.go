package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Product struct {
	Id      int
	StroeId int
	UserId  int
	Guid    string
	//产品头像与二维码
	ProductImage string

	//SEO
	Keywords    string
	Description string
	//产品基本信息
	ProductName    string //产品名称
	ProductGroupId int    //产品分组
	SerialNumber   string //序列号
	QRCode         string //二维码
	Barcode        string //条码
	RFID           string //RFID
	CustomID       string //自定义ID
	//产品参数
	Type          string //产品型号
	Color         string //产品颜色
	Size          string //产品尺寸
	Weight        int    //产品重量
	Unit          string //产品单位
	PCContent     string //PC详情内容
	MobileContent string //手机详情内容
	IsDefault     int    //是否默认产品
	//地区
	Province   string
	ProvinceId int
	City       string
	CityId     int
	Area       string
	AreaId     int
	//价格
	Price           float64 //零售价
	VIPPrice        float64 //会员价
	MarketPrice     float64 //市场价
	WholesalePrice  float64 //批发价
	WholesaleNumber int     //批发数量
	PurchasePrice   float64 //采购价

	//统计
	SaleCount        int //销售量
	SaleCommentCount int //评价数量
	ViewCount        int //页面访问数量
	//other
	Level int //搜索排序等级
	//0-未上架销售 1-上架销售
	OnSale int
	//0 -不显示  1 - 显示
	Status int
	//时间
	CreateTime int64
	ModifyTime int64
}

func (t *Product) TableName() string {
	return TableName("product")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Product))
}

func (t *Product) Update(fields ...string) error {
	if t.ProductName == "" {
		return fmt.Errorf("产品名称不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func ProductAdd(obj *Product) (int64, error) {
	if obj.ProductName == "" {
		return 0, fmt.Errorf("产品名称不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func ProductGetById(id int) (*Product, error) {
	obj := &Product{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func ProductDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("product")).Filter("id", id).Delete()
	return err
}

func ProductGetList(page, pageSize int) ([]*Product, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Product, 0)

	query := orm.NewOrm().QueryTable(TableName("product"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
