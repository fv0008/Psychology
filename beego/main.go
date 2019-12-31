package main

import (
	_ "../beego/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	maxIdle := 30
	maxConn := 30
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err := orm.RegisterDataBase("default", "sqlite3", "./datas/default.db", maxIdle, maxConn)
	//err := orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	if err != nil {
		fmt.Printf("数据库连接失败：%s\n", err.Error())
		return
	}
	// 开启日志记录
	logs.SetLogger("file")
	logs.SetLogger(logs.AdapterFile, `{"filename":"./log/run.log","maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
	logs.Async(1e3)

	if beego.BConfig.RunMode == "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	}
	beego.Run()
}
