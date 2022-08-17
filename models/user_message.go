package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"web/utils"
)

type User struct {
	Id         string
	Username   string
	Password   string
	Genre	   string // 1 普通用户， 2 管理员
	Status     int 	  // 0 正常状态， 1删除
	Createtime int64
}


//插入新用户数据到数据库
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,genre,status,createtime) values (?,?,?,?,?)",
		user.Username, user.Password, user.Genre, user.Status, user.Createtime)
}

//按条件查询用户,返回用户的id
func QueryUserWithCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	fmt.Println("user的id:",id)
	return id
}

//按条件查询用户收藏
func QueryUserCollectWithCon(con string) int {
	sql := fmt.Sprintf("select id from user_collect %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWithCon(sql)
}

//根据用户种类，查询id
func QueryUserWithGenre(username, password, genre string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s' and (genre = '%s')", username, password, genre)
	return QueryUserWithCon(sql)
}

//根据用户id，检查是否已经收藏
func QueryBookWithUserId(userId, bookId string) int {
	sql := fmt.Sprintf("where user_ids='%s' and novel_ids='%s'", userId, bookId)
	return QueryUserCollectWithCon(sql)
}

var usercileRowsNum  = 0

//只有首次获取行数的时候采取统计表里的行数,好像没啥用,前端处理好了
func GetUserRowsNum() int {
	if usercileRowsNum == 0 {
		usercileRowsNum = QueryNovelRowNum()
	}
	return usercileRowsNum
}

//设置用户页数
func SetUserRowsNum(){
	usercileRowsNum = QueryUserRowNum()
}

//查询用户的条数
func QueryUserRowNum() int {
	row := utils.QueryRowDB("select count(id) from users")
	num := 0
	row.Scan(&num)
	return num
}

//根据页码查询用户并返回数据以便展示
func FindUserWithPage(page,limit int) []orm.Params {
	orm.Debug = true

	page--
	sql := fmt.Sprintf("limit %d,%d", page*limit,limit)
	//fmt.Println("sql:::::::::::",sql)
	sql = "select id,username,genre from users " + sql
	//fmt.Println("sql:::::::::::",sql)

	var Users []orm.Params
	fmt.Println("var是正常的")
	i, e := db.Raw(sql).Values(&Users)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}

	return Users
}

//修改密码
func AlterPassword(username,password string) (int64, error){
	orm.Debug = true
	password = utils.MD5(password)
	sql := "update users SET password='" + password + "' "
	location := "where username='" + username + "'"
	sql = sql + location
	fmt.Println("sql::::",sql)
	return utils.ModifyDB(sql)
}

//验证旧密码是否输入一致
func TestOldPassword(oldPassword string) int {
	orm.Debug = true
	oldPassword = utils.MD5(oldPassword)
	sql := fmt.Sprintf("select id from users where password='%s'", oldPassword)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	fmt.Println("user的id:",id)
	return id
}