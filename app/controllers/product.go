package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
	"strconv"
	"strings"
)

type ProductController struct {
	BaseController
}

func (this *ProductController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.ProductGetList(page, this.pageSize)

	this.Data["pageTitle"] = "产品列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("ProductController.List"), true).ToString()
	this.display()
}

func (this *ProductController) Tree() {
	page := 1 //this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.ProductGetList(page, 2000)

	this.Data["pageTitle"] = "产品架构"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), 2000, beego.URLFor("ProductController.Tree"), true).ToString()
	this.display()
}

func (this *ProductController) Add() {
	if this.isPost() {
		product := new(models.Product)
		product.ProductName = strings.TrimSpace(this.GetString("product_name"))
		product.UserId = this.userId
		product.PCContent = strings.TrimSpace(this.GetString("description"))

		_, err := models.ProductAdd(product)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加产品"
	this.display()
}

func (this *ProductController) Edit() {
	id, _ := this.GetInt("id")

	product, err := models.ProductGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		product.ProductName = strings.TrimSpace(this.GetString("product_name"))
		product.PCContent = strings.TrimSpace(this.GetString("pc_conent"))
		product.MobileContent = strings.TrimSpace(this.GetString("mobile_conent"))
		err := product.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑产品"
	this.Data["product"] = product
	this.display()
}

func (this *ProductController) Batch() {
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
			models.ProductDelById(id)
			//models.TaskResetProductId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
