package models

import (
	"blog/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         string `orm:"column(uid);pk"` // 设置主键
	Username   string
	Password   string
	Salt	   string
	Phone	   string
	Img		   string
	Createtime int64
}

type Collect struct {
	Id		string
	UserId	string
	BlogId	string
}

type Blog struct {
	Id		string
	UserId  string
	Title	string
	Short	string
	Author	string
	Tag		string
	Content string
	Like	string
	Read	int
}

type Note struct {
	Id		string
	NotedId	string
	UserId 	string
	Content string
}

//插入新用户数据到数据库
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into blog_users(username,password,salt,img,phone,createtime) values (?,?,?,?,?,?)",
		user.Username, user.Password, user.Salt, user.Img, user.Phone, user.Createtime)
}
//插入新文章数据到数据库
func InsertBlog(blog Blog) (int64, error) {
	return utils.ModifyDB("insert into blog_info(user_id,blog_title,blog_short,blog_author,blog_tag,blog_content," +
		"blog_like,blog_read) values (?,?,?,?,?,?,?,?)",
		blog.UserId, blog.Title, blog.Short, blog.Author, blog.Tag, blog.Content, blog.Like, blog.Read)
}
//插入新好友数据到数据库
func InsertFriend(myId, f_Id string) (int64, error) {
	return utils.ModifyDB("insert into blog_friend(user_id,friend_id) values (?,?)", myId, f_Id)
}
//插入新留言数据到数据库
func InsertNote(note Note) (int64, error) {
	return utils.ModifyDB("insert into blog_note(noted_id,user_id,user_note) values (?,?,?)",
		note.NotedId,note.UserId,note.Content)
}
//更新新文章数据到数据库
func UpdateBlog(blog Blog) (int64, error) {
	return utils.ModifyDB("update blog_info set blog_title=?, blog_short=?, blog_tag=?, blog_content=? where id=?",
		blog.Title, blog.Short, blog.Tag, blog.Content, blog.Id)
}
//更新博客浏览量
func UpdateRead(id string) (int64, error) {
	return utils.ModifyDB("update blog_info set blog_read=blog_read+1 where id=?", id)
}
//点赞博客
func UpdateLike(id string) (int64, error) {
	return utils.ModifyDB("update blog_info set blog_like=blog_like+1 where id=?", id)
}
//更新用户信息
func UpdateInfo(id string, name string) (int64, error) {
	return utils.ModifyDB("update blog_users set username=? where id=?", name, id)
}
//更新用户头像
func UpdateImg(id string, img string) (int64, error) {
	return utils.ModifyDB("update blog_users set img=? where id=?", img, id)
}
//删除博客
func DeleteBlog(id string, userId string) (int64, error) {
	return utils.ModifyDB("delete from blog_info where id=? and user_id=?",id,userId)
}
//删除收藏
func DeleteCollect(userId, blogId string) (int64, error) {
	return utils.ModifyDB("delete from blog_collect where user_id=? and blog_id=?",userId,blogId)
}
//删除博客的同时删除收藏
func DeleteTogether(blogId string) (int64, error) {
	return utils.ModifyDB("delete from blog_collect where blog_id=?",blogId)
}
//插入新用户收藏数据到数据库
func InsertCollect(collect Collect) (int64, error) {
	return utils.ModifyDB("insert into blog_collect(user_id,blog_id) values (?,?)",
		collect.UserId, collect.BlogId)
}
//按条件查询用户,返回用户的id
func QueryUserWithCon(con string) int {
	sql := fmt.Sprintf("select id from blog_users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	fmt.Println("user的id:",id)
	return id
}
//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWithCon(sql)
}
//根据电话查询id
func QueryUserWithPhone(phone string) int {
	sql := fmt.Sprintf("where phone='%s'", phone)
	return QueryUserWithCon(sql)
}
//电话校验
func QueryUserWithPhoneAndId(phone, Id string) int {
	sql := fmt.Sprintf("where phone='%s' and id='%s'", phone, Id)
	return QueryUserWithCon(sql)
}
//返回用户的盐
func QuerySaltWithUsername(name string) string {
	sql := fmt.Sprintf("select salt from blog_users where username='%s'",name)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	salt := ""
	row.Scan(&salt)
	fmt.Println("user的盐:",salt)
	return salt
}
//返回用户名
func QueryUsernameWithId(Id string) string {
	sql := fmt.Sprintf("select username from blog_users where id='%s'",Id)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	username := ""
	row.Scan(&username)

	return username
}
//返回用户的头像
func QueryUserHeadImgWithUsername(name string) string{
	sql := fmt.Sprintf("select img from blog_users where username='%s'",name)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	img := ""
	row.Scan(&img)
	fmt.Println("user的头像:",img)
	return img
}
//返回用户的信息
func QueryUserInfoWithUsername(name string) []User{
	var user []User
	sql := fmt.Sprintf("select id,username,phone from blog_users where username='%s'",name)
	fmt.Println(sql)
	rows, _ := utils.QueryDB(sql)
	for rows.Next() {
		id := "0"
		username := ""
		phone := ""
		rows.Scan(&id,&username,&phone)
		tmp := User {
			Id: id, Username:username, Phone:phone,
		}
		user = append(user,tmp)
	}
	return user
}
//按条件查询用户收藏返回id
func QueryUserCollectWithCon(con string) int {
	sql := fmt.Sprintf("select id from blog_collect %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}
//根据id查询收藏
func QueryUserCollectWithId(user_Id int) []Blog {
	var blog []Blog
	sql := fmt.Sprintf("select * from blog_info where id in (select blog_id from blog_collect where user_id=%d)", user_Id)
	fmt.Println(sql)
	rows, _ := utils.QueryDB(sql)
	for rows.Next() {
		id := "0"
		userId := ""
		title := ""
		short := ""
		author := ""
		tag := ""
		content := ""
		like := ""
		read := 0
		rows.Scan(&id,&userId,&title,&short,&author,&tag,&content,&like,&read)
		tmp := Blog{
			Id: id, UserId: userId, Title: title, Short: short, Author: author,
			Tag: tag, Content: content, Like: like, Read: read,
		}
		blog = append(blog,tmp)
	}
	return blog
}
//根据id查询好友
func QueryUserFriendsWithId(user_Id int) []User {
	var friend []User
	sql := fmt.Sprintf("select id,username,phone from blog_users where id in " +
		"(select friend_id from blog_friend where user_id=%d)", user_Id)
	fmt.Println(sql)
	rows, _ := utils.QueryDB(sql)
	for rows.Next() {
		id := "0"
		userName := ""
		phone := ""
		rows.Scan(&id,&userName,&phone)
		tmp := User{
			Id: id, Username: userName, Phone: phone,
		}
		friend = append(friend,tmp)
	}
	return friend
}
//根据id查询博客内容
func QueryBlogWithBlogId(Id string) []Blog{
	var blog []Blog
	sql := fmt.Sprintf("select * from blog_info where id=%s", Id)
	fmt.Println(sql)
	rows, _ := utils.QueryDB(sql)
	for rows.Next() {
		id := "0"
		userId := ""
		title := ""
		short := ""
		author := ""
		tag := ""
		content := ""
		like := ""
		read := 0
		rows.Scan(&id,&userId,&title,&short,&author,&tag,&content,&like,&read)
		tmp := Blog{
			Id: id, UserId: userId, Title: title, Short: short, Author: author,
			Tag: tag, Content: content, Like: like, Read: read,
		}
		blog = append(blog,tmp)
	}
	return blog
}
//根据id查询留言内容
func QueryNoteWithNotedId(Id string) []Note{
	var note []Note
	sql := fmt.Sprintf("select * from blog_note where noted_id=%s", Id)
	fmt.Println(sql)
	rows, _ := utils.QueryDB(sql)
	for rows.Next() {
		id := "0"
		notedId := ""
		userId := ""
		content := ""
		rows.Scan(&id,&notedId,&userId,&content)
		tmp := Note{Id: id, NotedId: notedId, UserId: userId, Content: content}
		note = append(note,tmp)
	}
	return note
}
//根据用户id查询用户博客
func QueryBlogWithUserId(Id string) []Blog{
	var blog []Blog
	sql := fmt.Sprintf("select * from blog_info where user_id=%s", Id)
	fmt.Println(sql)
	rows, _ := utils.QueryDB(sql)
	for rows.Next() {
		id := "0"
		userId := ""
		title := ""
		short := ""
		author := ""
		tag := ""
		content := ""
		like := ""
		read := 0
		rows.Scan(&id,&userId,&title,&short,&author,&tag,&content,&like,&read)
		tmp := Blog{
			Id: id, UserId: userId, Title: title, Short: short, Author: author,
			Tag: tag, Content: content, Like: like, Read: read,
		}
		blog = append(blog,tmp)
	}
	return blog
}
//查询热搜榜前五
func QueryTop5Blog() []Blog{
	var blog []Blog
	sql := fmt.Sprintf("select * from blog_info order by blog_read desc limit 5")
	fmt.Println(sql)
	rows, _ := utils.QueryDB(sql)
	for rows.Next() {
		id := "0"
		userId := ""
		title := ""
		short := ""
		author := ""
		tag := ""
		content := ""
		like := ""
		read := 0
		rows.Scan(&id,&userId,&title,&short,&author,&tag,&content,&like,&read)
		tmp := Blog{
			Id: id, UserId: userId, Title: title, Short: short, Author: author,
			Tag: tag, Content: content, Like: like, Read: read,
		}
		blog = append(blog,tmp)
	}
	return blog
}
//修改密码
func AlterPassword(username,password string) (int64, error){
	orm.Debug = true
	password = utils.MD5(password)
	salt := utils.MD5(utils.Salt())
	sql := "update blog_users SET password='" + password + salt + "', salt='" + salt + "' "
	location := "where username='" + username + "'"
	sql = sql + location
	fmt.Println("sql::::",sql)
	return utils.ModifyDB(sql)
}

//验证旧密码是否输入一致
func TestOldPassword(username string, oldPassword string) int {
	orm.Debug = true
	oldPassword = utils.MD5(oldPassword)
	sql := fmt.Sprintf("select id from blog_users where password='%s'", oldPassword + QuerySaltWithUsername(username))
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	fmt.Println("user的id:",id)
	return id
}