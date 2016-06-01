package main

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/wqdsoft/moilicms/app/jobs"
	"github.com/wqdsoft/moilicms/app/models"
	_ "github.com/wqdsoft/moilicms/app/routers"
	"html/template"
	"net/http"
	"os"
	"strings"
)

const VERSION = "1.0.0"

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-install" {
			fmt.Println("System will be clear all tables and rebuild all tables,Are you sure?(y/yes)")
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()
			command := string(data)
			command = strings.ToLower(command)
			if command == "yes" || command == "y" {
				models.Syncdb()
				os.Exit(0)
			}
		}
	}
}

func main() {
	initArgs()
	//判断是否安装
	models.Init()
	jobs.InitJobs()
	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.AppConfig.Set("version", VERSION)
	beego.Run()
}
