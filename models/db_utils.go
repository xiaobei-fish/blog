package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db orm.Ormer
)

func init() {
	orm.Debug = true //是否开启调试模式，调试模式下会打印SQL语句

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	_ = orm.RegisterDataBase("default", "mysql", dbConn)
	orm.RegisterModel(new(User))
	var tmp = orm.NewOrm()
	db = tmp
}
