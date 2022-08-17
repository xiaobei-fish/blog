package main

import (
	"github.com/astaxie/beego"
	_ "web/routers"
	"web/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

