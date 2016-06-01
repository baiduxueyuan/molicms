package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
)

type FinancialController struct {
	BaseController
}

func (this *FinancialController) RefundList() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.RefundGetList(page, this.pageSize)

	this.Data["pageTitle"] = "退款审核"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("FinancialController.List"), true).ToString()
	this.display()
}
