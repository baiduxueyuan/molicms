package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
	"strconv"
	"strings"
)

type InventoryController struct {
	BaseController
}

func (this *InventoryController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.InventoryGetList(page, this.pageSize)

	this.Data["pageTitle"] = "库存列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("InventoryController.List"), true).ToString()
	this.display()
}

func (this *InventoryController) Add() {
	if this.isPost() {
		inventory := new(models.Inventory)
		inventory.ProductName = strings.TrimSpace(this.GetString("product_name"))
		inventory.UserId = this.userId

		_, err := models.InventoryAdd(inventory)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加库存"
	this.display()
}

func (this *InventoryController) Edit() {
	id, _ := this.GetInt("id")

	inventory, err := models.InventoryGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		inventory.ProductName = strings.TrimSpace(this.GetString("product_name"))

		err := inventory.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑库存"
	this.Data["inventory"] = inventory
	this.display()
}

func (this *InventoryController) Batch() {
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
			models.InventoryDelById(id)
			//models.TaskResetInventoryId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
