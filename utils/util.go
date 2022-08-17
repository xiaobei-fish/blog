package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {

	fmt.Println("InitMysql....")
	driverName := beego.AppConfig.String("driverName")

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
		CreateTableNamedByUserCollect()
	}
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	//fmt.Println("sql::",sql)
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		genre INT(4),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

//创建用书收藏夹表
func CreateTableNamedByUserCollect() {
	sql := `CREATE TABLE IF NOT EXISTS user_collect(
		id INT(11)  NOT NULL PRIMARY KEY AUTO_INCREMENT,
		novel_ids INT(4) NOT NULL,
		user_ids INT(10),
		CONSTRAINT fk_user_collect_novel_info FOREIGN KEY (novel_ids) REFERENCES novel_info (id)  on delete restrict   ON UPDATE CASCADE
		)AUTO_INCREMENT = 1;`
	ModifyDB(sql)
}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

//MD5加密密码
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}
