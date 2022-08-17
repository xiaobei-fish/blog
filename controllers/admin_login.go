package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"web/models"
	"web/utils"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.TplName = "home.html"
}

func (c *AdminController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("username:", username, ",password:", password)

	//通过genre限制了权限,只有管理员才可以登录
	id := models.QueryUserWithGenre(username, utils.MD5(password), "2")
	fmt.Println("id:", id)

	if id > 0 {
		//设置session以免密登录
		c.SetSession("loginuser", username)
		c.SetSession("userid",id)

		c.Data["json"] = map[string]interface{}{"code": 1, "message": "成功   XX小说网欢迎您"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "失败,密码错误或您不是管理员"}
	}
	c.ServeJSON()
}
