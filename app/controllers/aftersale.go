package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
	"strconv"
	"strings"
)

type AftersaleController struct {
	BaseController
}

func (this *AftersaleController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.AftersaleGetList(page, this.pageSize)

	this.Data["pageTitle"] = "售后服务列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("AftersaleController.List"), true).ToString()
	this.display()
}

func (this *AftersaleController) Add() {
	if this.isPost() {
		aftersale := new(models.Aftersale)
		aftersale.AftersaleName = strings.TrimSpace(this.GetString("aftersale_name"))
		aftersale.UserId = this.userId
		aftersale.Remark = strings.TrimSpace(this.GetString("remark"))

		_, err := models.AftersaleAdd(aftersale)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加售后服务"
	this.display()
}

func (this *AftersaleController) Edit() {
	id, _ := this.GetInt("id")

	aftersale, err := models.AftersaleGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		aftersale.AftersaleName = strings.TrimSpace(this.GetString("aftersale_name"))
		aftersale.Remark = strings.TrimSpace(this.GetString("remark"))
		err := aftersale.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑售后服务"
	this.Data["aftersale"] = aftersale
	this.display()
}

func (this *AftersaleController) Batch() {
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
			models.AftersaleDelById(id)
			//models.TaskResetAftersaleId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
