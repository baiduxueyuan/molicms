package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
	"strconv"
	"strings"
)

type ContentController struct {
	BaseController
}

func (this *ContentController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.ContentGetList(page, this.pageSize)

	this.Data["pageTitle"] = "素材列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("ContentController.List"), true).ToString()
	this.display()
}

func (this *ContentController) Tree() {
	page := 1 //this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.ContentGetList(page, 2000)

	this.Data["pageTitle"] = "素材架构"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), 2000, beego.URLFor("ContentController.Tree"), true).ToString()
	this.display()
}

func (this *ContentController) Add() {
	if this.isPost() {
		content := new(models.Content)
		content.ContentName = strings.TrimSpace(this.GetString("content_name"))
		content.UserId = this.userId
		content.Description = strings.TrimSpace(this.GetString("description"))

		_, err := models.ContentAdd(content)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加素材"
	this.display()
}

func (this *ContentController) Edit() {
	id, _ := this.GetInt("id")

	content, err := models.ContentGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		content.ContentName = strings.TrimSpace(this.GetString("content_name"))
		content.Description = strings.TrimSpace(this.GetString("description"))
		err := content.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑素材"
	this.Data["content"] = content
	this.display()
}

func (this *ContentController) Batch() {
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
			models.ContentDelById(id)
			//models.TaskResetContentId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
