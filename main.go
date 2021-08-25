package main

import (
	_ "blog/routers"
	"blog/utils"
	"github.com/astaxie/beego"
)


func main() {

	utils.InitMysql()
	beego.Run()
}

