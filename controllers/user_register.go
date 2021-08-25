package controllers

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "home.html"
}

//处理注册
func (c *RegisterController) Post() {

	//获取表单信息
	username := c.GetString("username")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	phone := c.GetString("phone")
	fmt.Printf("username:%s,password:%s,repassword:%s,phone:%s\n",username,password,repassword,phone)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id0 := models.QueryUserWithUsername(username)
	id1 := models.QueryUserWithPhone(phone)
	fmt.Println("id0:", id0)
	fmt.Println("id1:", id1)
	if id0 > 0 {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "该用户名已经存在"}

		c.ServeJSON()
		return
	}else if id1 > 0 {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "该邮箱已经被注册过"}

		c.ServeJSON()
		return
	}

	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	fmt.Println("md5后：", password)
	//储存的密码： MD5密码+MD5盐
	user := models.User{Username: username, Password: password + utils.MD5(utils.Salt()), Salt: utils.MD5(utils.Salt()),
		Createtime: time.Now().Unix(), Phone: phone , Img: "../static/img/default.jpg"}

	_, err := models.InsertUser(user)

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	c.ServeJSON()
}