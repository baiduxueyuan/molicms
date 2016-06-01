package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
	"strconv"
	"strings"
)

type OrderController struct {
	BaseController
}

func (this *OrderController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.OrderGetList(page, this.pageSize)

	this.Data["pageTitle"] = "订单列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("OrderController.List"), true).ToString()
	this.display()
}

func (this *OrderController) Tree() {
	page := 1 //this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.OrderGetList(page, 2000)

	this.Data["pageTitle"] = "订单架构"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), 2000, beego.URLFor("OrderController.Tree"), true).ToString()
	this.display()
}

func (this *OrderController) Add() {
	if this.isPost() {
		order := new(models.Order)
		order.OrderName = strings.TrimSpace(this.GetString("order_name"))
		order.UserId = this.userId

		_, err := models.OrderAdd(order)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加订单"
	this.display()
}

func (this *OrderController) Edit() {
	id, _ := this.GetInt("id")

	order, err := models.OrderGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		order.OrderName = strings.TrimSpace(this.GetString("order_name"))

		err := order.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑订单"
	this.Data["order"] = order
	this.display()
}

func (this *OrderController) Batch() {
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
			models.OrderDelById(id)
			//models.TaskResetOrderId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
