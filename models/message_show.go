package models

import (
	"blog/utils"
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)
//搜索
type SearchBlockParam struct {
	Id         string
	Title      string
	Tag        string
	Short      string
	Author     string
	Like 	   string
	Read	   int

	//查看文章的地址
	Link string

	//记录是否登录
	IsLogin bool
}
//主页
type HomeBlockParam struct {
	Id      string
	Title	string
	Author	string
	Tag		string
	Read	int
}
//分页结构体
type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}
//个人主页
type SelfBlockParam struct {
	Id         string
	Title      string
	Tags       string
	Short      string
	Like 	   string
	Read	   int

	//查看文章的地址
	Link string
}
//好友主页
type FriendBlockParam struct {
	Id         string
	Username   string
	Phone      string
}
//留言板块
type NoteBlockParam struct {
	Id        string
	NotedId   string
	UserId    string
	Content	  string
}
//首页模板
func MakeHomeBlocks(blogs []Blog) template.HTML {
	htmlHome := ""
	for _, blog := range blogs {
		fmt.Println(blog)
		homeParam := HomeBlockParam{}
		homeParam.Id = blog.Id
		homeParam.Title = blog.Title
		homeParam.Author = blog.Author
		homeParam.Tag = blog.Tag
		homeParam.Read = blog.Read

		fmt.Printf("title:%s,author:%s,tag:%s,read:%d,id:%s\n",blog.Title,blog.Author,blog.Tag,blog.Read,blog.Id)

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block3.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}
//搜索模板
func MakeSearchBlocks(blogs []Blog) template.HTML {
	htmlHome := ""
	for _, blog := range blogs {
		homeParam := SearchBlockParam{}

		homeParam.Id = blog.Id
		homeParam.Title = blog.Title
		homeParam.Author = blog.Author
		homeParam.Short = blog.Short
		homeParam.Tag = blog.Tag
		homeParam.Like = blog.Like
		homeParam.Read = blog.Read

		homeParam.Link = "/detail?id=" + blog.Id

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}
//我的博客模板
func MakeMineBlocks(blogs []Blog) template.HTML {
	htmlHome := ""
	for _, blog := range blogs {
		fmt.Println(blog)
		homeParam := SearchBlockParam{}

		homeParam.Id = blog.Id
		homeParam.Title = blog.Title
		homeParam.Author = blog.Author
		homeParam.Short = blog.Short
		homeParam.Tag = blog.Tag
		homeParam.Like = blog.Like
		homeParam.Read = blog.Read

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block4.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}
//我的收藏模板
func MakeCollectBlocks(blogs []Blog) template.HTML {
	htmlHome := ""
	for _, blog := range blogs {
		fmt.Println(blog)
		homeParam := SearchBlockParam{}

		homeParam.Id = blog.Id
		homeParam.Title = blog.Title
		homeParam.Author = blog.Author
		homeParam.Short = blog.Short
		homeParam.Tag = blog.Tag
		homeParam.Like = blog.Like
		homeParam.Read = blog.Read

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block2.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}
//好友页面模板
func MakeFriendBlocks(users []User) template.HTML {
	htmlHome := ""
	for _, user := range users {
		fmt.Println(user)
		homeParam := FriendBlockParam{}
		homeParam.Id = user.Id
		homeParam.Username = user.Username
		homeParam.Phone = user.Phone

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block5.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}
//留言小版块模板
func MakeSmallNoteBlocks(notes []Note) template.HTML {
	htmlHome := ""
	for _, note := range notes {
		fmt.Println(note)
		homeParam := NoteBlockParam{}
		homeParam.Id = note.Id
		homeParam.NotedId = note.NotedId
		homeParam.UserId = note.UserId
		if len(note.Content) > 10 {
			homeParam.Content = note.Content[0: 10]
		}else{
			homeParam.Content = note.Content
		}

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block6.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}
//留言版块模板
func MakeNoteBlocks(notes []Note) template.HTML {
	htmlHome := ""
	for _, note := range notes {
		fmt.Println(note)
		homeParam := NoteBlockParam{}
		homeParam.Id = note.Id
		homeParam.NotedId = note.NotedId
		homeParam.UserId = note.UserId
		homeParam.Content = note.Content

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block6.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}
//翻页,page是当前的页数
func ConfigHomeFooterPageCode(page int, key string) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	//查询出总的条数
	num := GetBlogRowsNum()
	//从配置文件中读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("blogListPageNum")
	//计算出总页数
	allPageNum := (num-1)/pageRow + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1，那么上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/search?key=" + key + "&page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/search?key=" + key + "&page=" + strconv.Itoa(page+1)
	return pageCode
}
//翻页,page是当前的页数
func ConfigAllHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	//查询出总的条数
	num := GetBlogRowsNum()
	//从配置文件中读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("blogListPageNum")
	//计算出总页数
	allPageNum := (num-1)/pageRow + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1，那么上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/all?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/all?page=" + strconv.Itoa(page+1)
	return pageCode
}
//根据页码查询全部文章
func FindAllBlogWithPage(page int) ([]Blog, error) {
	//从配置文件中获取每页的文章数量
	num, _ := beego.AppConfig.Int("blogListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryAllBlogWithPage(page, num)
}
//根据页码查询文章
func FindBlogWithPage(page int, key string) ([]Blog, error) {
	//从配置文件中获取每页的文章数量
	num, _ := beego.AppConfig.Int("blogListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryBlogWithPage(page, num, key)
}
//分页查询全部
func QueryAllBlogWithPage(page, num int) ([]Blog, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryBlogsWithCon(sql)
}
//分页查询
func QueryBlogWithPage(page, num int, key string) ([]Blog, error) {
	con := "'%" + key + "%' "
	Title := "blog_title like " + con
	Short := "blog_short like " + con
	Author := "blog_author like " + con
	Tag := "blog_tag like " + con
	Content := "blog_content like " + con

	where := "where " + Title + "or " + Short + "or " + Author + "or " + Tag + "or " + Content

	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryBlogsWithCon(where + sql)
}
func QueryBlogsWithCon(sql string) ([]Blog, error) {
	sql = "select * from blog_info " + sql
	fmt.Println("sql:",sql)
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var blog []Blog
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
	return blog, nil
}
//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var blogRowsNum = 0
//只有首次获取行数的时候采取统计表里的行数
func GetBlogRowsNum() int {
	if blogRowsNum == 0 {
		blogRowsNum = QueryBlogRowNum()
	}
	return blogRowsNum
}
//查询文章的总条数
func QueryBlogRowNum() int {
	row := utils.QueryRowDB("select count(id) from blog_info")
	num := 0
	row.Scan(&num)
	return num
}
//设置页数,更新总博客数目
func SetBlogRowsNum(){
	blogRowsNum = QueryBlogRowNum()
}
