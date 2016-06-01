package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
	"strconv"
	"strings"
)

type CustomerController struct {
	BaseController
}

func (this *CustomerController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.CustomerGetList(page, this.pageSize)

	this.Data["pageTitle"] = "客户列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("CustomerController.List"), true).ToString()
	this.display()
}

func (this *CustomerController) Add() {
	if this.isPost() {
		customer := new(models.Customer)
		customer.CustomerName = strings.TrimSpace(this.GetString("customer_name"))
		customer.UserId = this.userId
		customer.Remark = strings.TrimSpace(this.GetString("remark"))

		_, err := models.CustomerAdd(customer)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加客户"
	this.display()
}

func (this *CustomerController) Edit() {
	id, _ := this.GetInt("id")

	customer, err := models.CustomerGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		customer.CustomerName = strings.TrimSpace(this.GetString("customer_name"))
		customer.Remark = strings.TrimSpace(this.GetString("remark"))
		err := customer.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑客户"
	this.Data["customer"] = customer
	this.display()
}

func (this *CustomerController) Batch() {
	action := this.GetString("action")
	ids := this.GetStrings("ids")
	if len(ids) < 1 {
		this.ajaxMsg("请选择要操作的项目", MSG_ERR)
	}

	for _, v := range ids {
		id, _ := strconv.Atoi(v)
		if id < 1 {
			continue
		}
		switch action {
		case "delete":
			models.CustomerDelById(id)
			//models.TaskResetCustomerId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
