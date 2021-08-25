package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/astaxie/beego"
	"log"
	"time"
)

var db *sql.DB

func InitMysql(){
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
		CreateTablbWithBlog()
		CreateTablbWithCollect()
		CreateTablbWithFriends()
		CreateTablbWithNote()
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
	sql := `CREATE TABLE IF NOT EXISTS blog_users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		salt VARCHAR(64),
		img VARCHAR(64),
		phone VARCHAR(64),
		createtime INT(10)
		);`
	ModifyDB(sql)
}
//创建博客信息表
func CreateTablbWithBlog()  {
	sql := `CREATE TABLE IF NOT EXISTS blog_info(
		id INT(11)  NOT NULL PRIMARY KEY AUTO_INCREMENT,
		user_id INT(10) NOT NULL,
		blog_title VARCHAR(64),
		blog_short VARCHAR(255),
		blog_author VARCHAR(64),
		blog_tag VARCHAR(64),
		blog_content longtext,
		blog_like INT(10),
		blog_read INT(10)
		);`
	ModifyDB(sql)
}
//创建用户收藏表
func CreateTablbWithCollect()  {
	sql := `CREATE TABLE IF NOT EXISTS blog_collect(
		id INT(11)  NOT NULL PRIMARY KEY AUTO_INCREMENT,
		user_id INT(10) NOT NULL,
		blog_id INT(10) NOT NULL
		);`
	ModifyDB(sql)
}
//创建用户好友表
func CreateTablbWithFriends()  {
	sql := `CREATE TABLE IF NOT EXISTS blog_friend(
		id INT(11)  NOT NULL PRIMARY KEY AUTO_INCREMENT,
		user_id INT(10) NOT NULL,
		friend_id INT(10) NOT NULL
		);`
	ModifyDB(sql)
}
//创建用户留言表
func CreateTablbWithNote()  {
	sql := `CREATE TABLE IF NOT EXISTS blog_note(
		id INT(11)  NOT NULL PRIMARY KEY AUTO_INCREMENT,
		noted_id INT(10) NOT NULL,
		user_id INT(10) NOT NULL,
		user_note longtext
		);`
	ModifyDB(sql)
}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
//MD5加密密码
func MD5(str string) string{
	md5str := fmt.Sprintf("%x",md5.Sum([]byte(str)))
	return  md5str
}
//加盐
func Salt() string{
	salt := time.Now().Unix()
	return string(salt)
}