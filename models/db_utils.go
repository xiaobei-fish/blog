package models

import (
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/astaxie/beego/orm"
)

var (
	db orm.Ormer
)

func init(){
	orm.Debug = true //是否开启调试模式，调试模式下会打印SQL语句
	_ = orm.RegisterDataBase("default", "mysql", "root:qwe123@tcp(127.0.0.1:3306)/west?charset=utf8")
	orm.RegisterModel(new(User))
	var tmp = orm.NewOrm()
	db = tmp
}
