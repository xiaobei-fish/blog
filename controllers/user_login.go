package controllers

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "home.html"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("username:", username, ",password:", password)

	con0 := fmt.Sprintf("where username='%s' and password='%s'",username,utils.MD5(password) + models.QuerySaltWithUsername(username))
	con1 := fmt.Sprintf("where phone='%s' and password='%s'",username,utils.MD5(password) + models.QuerySaltWithUsername(username))

	id0 := models.QueryUserWithCon(con0)
	id1 := models.QueryUserWithCon(con1)
	fmt.Println("id0:", id0)
	fmt.Println("id1:",id1)

	if id0 > 0 || id1 > 0 {
		//设置session以免密登录
		c.SetSession("loginuser", username)

		c.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	c.ServeJSON()
}
