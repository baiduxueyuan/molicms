package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
	"strconv"
	"strings"
)

type WarehouseController struct {
	BaseController
}

func (this *WarehouseController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.WarehouseGetList(page, this.pageSize)

	this.Data["pageTitle"] = "仓库列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("WarehouseController.List"), true).ToString()
	this.display()
}

func (this *WarehouseController) LocationList() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.WarehouseGetList(page, this.pageSize)

	this.Data["pageTitle"] = "库位列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("WarehouseController.LocationList"), true).ToString()
	this.display()
}

func (this *WarehouseController) AreaList() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.WarehouseGetList(page, this.pageSize)

	this.Data["pageTitle"] = "仓区列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("WarehouseController.AreaList"), true).ToString()
	this.display()
}

func (this *WarehouseController) Add() {
	if this.isPost() {
		warehouse := new(models.Warehouse)
		warehouse.WarehouseName = strings.TrimSpace(this.GetString("warehouse_name"))
		warehouse.UserId = this.userId

		_, err := models.WarehouseAdd(warehouse)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加仓库"
	this.display()
}

func (this *WarehouseController) AreaAdd() {
	if this.isPost() {
		warehouse := new(models.WarehouseArea)
		warehouse.WarehouseAreaName = strings.TrimSpace(this.GetString("warehouse_area_name"))
		warehouse.UserId = this.userId

		_, err := models.WarehouseAreaAdd(warehouse)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加仓区"
	this.display()
}

func (this *WarehouseController) LocationAdd() {
	if this.isPost() {
		warehouse := new(models.WarehouseLocation)
		warehouse.WarehouseLocationName = strings.TrimSpace(this.GetString("warehouse_location_name"))
		warehouse.UserId = this.userId

		_, err := models.WarehouseLocationAdd(warehouse)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加库位"
	this.display()
}

func (this *WarehouseController) Edit() {
	id, _ := this.GetInt("id")

	warehouse, err := models.WarehouseGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		warehouse.WarehouseName = strings.TrimSpace(this.GetString("warehouse_name"))

		err := warehouse.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑仓库"
	this.Data["warehouse"] = warehouse
	this.display()
}

func (this *WarehouseController) LocationEdit() {
	id, _ := this.GetInt("id")

	warehouse, err := models.WarehouseLocationGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		warehouse.WarehouseLocationName = strings.TrimSpace(this.GetString("warehouse_location_name"))

		err := warehouse.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑库位"
	this.Data["warehouse"] = warehouse
	this.display()
}

func (this *WarehouseController) AreaEdit() {
	id, _ := this.GetInt("id")

	warehouse, err := models.WarehouseAreaGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		warehouse.WarehouseAreaName = strings.TrimSpace(this.GetString("warehouse_area_name"))

		err := warehouse.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑仓区"
	this.Data["warehouse"] = warehouse
	this.display()
}

func (this *WarehouseController) Batch() {
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
			models.WarehouseDelById(id)
			//models.TaskResetWarehouseId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}

func (this *WarehouseController) LocationBatch() {
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
			models.WarehouseLocationDelById(id)
			//models.TaskResetWarehouseId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}

func (this *WarehouseController) AreaBatch() {
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
			models.WarehouseAreaDelById(id)
			//models.TaskResetWarehouseId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
