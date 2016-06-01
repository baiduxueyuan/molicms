package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
	"strconv"
	"strings"
)

type AgentController struct {
	BaseController
}

func (this *AgentController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.AgentGetList(page, this.pageSize)

	this.Data["pageTitle"] = "代理商列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("AgentController.List"), true).ToString()
	this.display()
}

func (this *AgentController) Tree() {
	page := 1 //this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := models.AgentGetList(page, 2000)

	this.Data["pageTitle"] = "代理商架构"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), 2000, beego.URLFor("AgentController.Tree"), true).ToString()
	this.display()
}

func (this *AgentController) Add() {
	if this.isPost() {
		agent := new(models.Agent)
		agent.AgentName = strings.TrimSpace(this.GetString("agent_name"))
		agent.UserId = this.userId
		agent.Description = strings.TrimSpace(this.GetString("description"))

		_, err := models.AgentAdd(agent)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加代理商"
	this.display()
}

func (this *AgentController) Edit() {
	id, _ := this.GetInt("id")

	agent, err := models.AgentGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		agent.AgentName = strings.TrimSpace(this.GetString("agent_name"))
		agent.Description = strings.TrimSpace(this.GetString("description"))
		err := agent.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑代理商"
	this.Data["agent"] = agent
	this.display()
}

func (this *AgentController) Batch() {
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
			models.AgentDelById(id)
			//models.TaskResetAgentId(id)
		}
	}

	this.ajaxMsg("", MSG_OK)
}
