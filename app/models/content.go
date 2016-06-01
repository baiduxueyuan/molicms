package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Content struct {
	Id             int
	OrganizationId int
	UserId         int
	ContentName    string
	ParentId       int
	RealName       string
	Weixin         string
	Qq             string
	Mobile         string
	IdCard         string
	Province       string
	ProvinceId     int
	City           string
	CityId         int
	Area           string
	AreaId         int
	Address        string
	Description    string
	Level          int
	Status         int
	CreateTime     int64
}

func (t *Content) TableName() string {
	return TableName("content")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Content))
}

func (t *Content) Update(fields ...string) error {
	if t.ContentName == "" {
		return fmt.Errorf("代理商名称不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func ContentAdd(obj *Content) (int64, error) {
	if obj.ContentName == "" {
		return 0, fmt.Errorf("代理商名称不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func ContentGetById(id int) (*Content, error) {
	obj := &Content{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func ContentDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("content")).Filter("id", id).Delete()
	return err
}

func ContentGetList(page, pageSize int) ([]*Content, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Content, 0)

	query := orm.NewOrm().QueryTable(TableName("content"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
