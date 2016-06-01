package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/jobs"
	"github.com/wqdsoft/moilicms/app/libs"
	"github.com/wqdsoft/moilicms/app/models"
	"runtime"
	"time"
)

type InstallController struct {
	BaseController
}

// 首页
func (this *InstallController) Index() {
	this.Data["pageTitle"] = "系统概况"

	// 即将执行的任务
	entries := jobs.GetEntries(30)
	jobList := make([]map[string]interface{}, len(entries))
	for k, v := range entries {
		row := make(map[string]interface{})
		job := v.Job.(*jobs.Job)
		row["task_id"] = job.GetId()
		row["task_name"] = job.GetName()
		row["next_time"] = beego.Date(v.Next, "Y-m-d H:i:s")
		jobList[k] = row
	}

	// 最近执行的日志
	logs, _ := models.TaskLogGetList(1, 20)
	recentLogs := make([]map[string]interface{}, len(logs))
	for k, v := range logs {
		task, err := models.TaskGetById(v.TaskId)
		taskName := ""
		if err == nil {
			taskName = task.TaskName
		}
		row := make(map[string]interface{})
		row["task_name"] = taskName
		row["id"] = v.Id
		row["start_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["process_time"] = float64(v.ProcessTime) / 1000
		row["ouput_size"] = libs.SizeFormat(float64(len(v.Output)))
		row["output"] = beego.Substr(v.Output, 0, 100)
		row["status"] = v.Status
		recentLogs[k] = row
	}

	// 最近执行失败的日志
	logs, _ = models.TaskLogGetList(1, 20, "status__lt", 0)
	errLogs := make([]map[string]interface{}, len(logs))
	for k, v := range logs {
		task, err := models.TaskGetById(v.TaskId)
		taskName := ""
		if err == nil {
			taskName = task.TaskName
		}
		row := make(map[string]interface{})
		row["task_name"] = taskName
		row["id"] = v.Id
		row["start_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["process_time"] = float64(v.ProcessTime) / 1000
		row["ouput_size"] = libs.SizeFormat(float64(len(v.Output)))
		row["error"] = beego.Substr(v.Error, 0, 100)
		row["status"] = v.Status
		errLogs[k] = row
	}

	this.Data["recentLogs"] = recentLogs
	this.Data["errLogs"] = errLogs
	this.Data["jobs"] = jobList
	this.Data["cpuNum"] = runtime.NumCPU()
	this.TplName = "install/index.html"
}

// 获取系统时间
func (this *InstallController) GetTime() {
	out := make(map[string]interface{})
	out["time"] = time.Now().UnixNano() / int64(time.Millisecond)
	this.jsonResult(out)
}
