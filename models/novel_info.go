package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db orm.Ormer
)

type NovelInfo struct {
	Id                 int64
	Novel_id           string
	Novel_name         string
	Novel_pic          string
	Novel_writer       string
	Novel_introduction string
	Novel_genre        string
	Novel_state        string
}

func init() {
	orm.Debug = true //是否开启调试模式，调试模式下会打印SQL语句
	orm.RegisterDataBase("default", "mysql", "root:qwe123@tcp(127.0.0.1:3306)/bottle?charset=utf8")
	orm.RegisterModel(new(NovelInfo))
	db = orm.NewOrm()
}
func AddNovels(novelInfo *NovelInfo) (int64, error) {
	id, err := db.Insert(novelInfo)
	return id, err
}
