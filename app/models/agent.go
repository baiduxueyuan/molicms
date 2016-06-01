package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Agent struct {
	Id             int
	OrganizationId int
	UserId         int
	AgentName      string
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

func (t *Agent) TableName() string {
	return TableName("agent")
}

//注册model实现自动生成表
func init() {
	orm.RegisterModel(new(Agent))
}

func (t *Agent) Update(fields ...string) error {
	if t.AgentName == "" {
		return fmt.Errorf("代理商名称不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func AgentAdd(obj *Agent) (int64, error) {
	if obj.AgentName == "" {
		return 0, fmt.Errorf("代理商名称不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func AgentGetById(id int) (*Agent, error) {
	obj := &Agent{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func AgentDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("agent")).Filter("id", id).Delete()
	return err
}

func AgentGetList(page, pageSize int) ([]*Agent, int64) {
	offset := (page - 1) * pageSize

	list := make([]*Agent, 0)

	query := orm.NewOrm().QueryTable(TableName("agent"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}
