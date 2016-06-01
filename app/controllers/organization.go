package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
	"strconv"
	"strings"
)

type OrganizationController struct {
	BaseController
}

func (this *OrganizationController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.OrganizationGetList(page, this.pageSize)

	this.Data["pageTitle"] = "组织列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("OrganizationController.List"), true).ToString()
	this.display()
}

func (this *OrganizationController) Tree() {
	page := 1 //this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.OrganizationGetList(page, 2000)

	this.Data["pageTitle"] = "组织架构"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), 2000, beego.URLFor("OrganizationController.Tree"), true).ToString()
	this.display()
}

func (this *OrganizationController) Add() {
	if this.isPost() {
		organization := new(models.Organization)
		organization.OrganizationName = strings.TrimSpace(this.GetString("organization_name"))
		organization.UserId = this.userId
		organization.Description = strings.TrimSpace(this.GetString("description"))

		_, err := models.OrganizationAdd(organization)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加组织"
	this.display()
}

func (this *OrganizationController) Edit() {
	id, _ := this.GetInt("id")

	organization, err := models.OrganizationGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		organization.OrganizationName = strings.TrimSpace(this.GetString("organization_name"))
		organization.Description = strings.TrimSpace(this.GetString("description"))
		err := organization.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑组织"
	this.Data["organization"] = organization
	this.display()
}

func (this *OrganizationController) Batch() {
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
			models.OrganizationDelById(id)
			//models.TaskResetOrganizationId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
