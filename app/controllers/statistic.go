package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
)

type StatisticController struct {
	BaseController
}

func (this *StatisticController) Overview() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.OrderGetList(page, this.pageSize)

	this.Data["pageTitle"] = "订单统计"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("StatisticController.List"), true).ToString()
	this.display()
}

func (this *StatisticController) Orderstat() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.OrderGetList(page, this.pageSize)

	this.Data["pageTitle"] = "订单统计"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("StatisticController.List"), true).ToString()
	this.display()
}

func (this *StatisticController) Productstat() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.ProductGetList(page, this.pageSize)

	this.Data["pageTitle"] = "商品统计"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("StatisticController.List"), true).ToString()
	this.display()
}

func (this *StatisticController) Customerstat() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.CustomerGetList(page, this.pageSize)

	this.Data["pageTitle"] = "客户统计"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("StatisticController.List"), true).ToString()
	this.display()
}

func (this *StatisticController) Contentstat() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.ContentGetList(page, this.pageSize)

	this.Data["pageTitle"] = "转发统计"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("StatisticController.List"), true).ToString()
	this.display()
}
