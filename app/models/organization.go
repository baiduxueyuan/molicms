package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Organization struct {
	Id               int
	OrganizationId   int
	UserId           int
	OrganizationName string
	Qq               string
	Mobile           string
	Province         string
	ProvinceId       int
	City             string
	CityId           int
	Area             string
	AreaId           int
	Address          string
	Description      string
	Level            int
	Status           int
	CreateTime       int64
}

func (t *Organization) TableName() string {
	return TableName("organization")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Organization))
}

func (t *Organization) Update(fields ...string) error {
	if t.OrganizationName == "" {
		return fmt.Errorf("组织名称不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func OrganizationAdd(obj *Organization) (int64, error) {
	if obj.OrganizationName == "" {
		return 0, fmt.Errorf("组织名称不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func OrganizationGetById(id int) (*Organization, error) {
	obj := &Organization{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func OrganizationDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("organization")).Filter("id", id).Delete()
	return err
}

func OrganizationGetList(page, pageSize int) ([]*Organization, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Organization, 0)

	query := orm.NewOrm().QueryTable(TableName("organization"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
