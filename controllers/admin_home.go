package controllers

import (
	"github.com/astaxie/goredis"
)

type AdminHomeController struct {
	JudgeController
}

type Picture struct {
	Url	string
}

func (c *AdminHomeController) Get() {
	c.Data["adminName"] = c.Loginuser
	c.TplName = "admin_home.html"
}

func (c *AdminHomeController) Post() {
	id1   := c.GetString("slideshow_1")
	id2   := c.GetString("slideshow_2")
	id3   := c.GetString("slideshow_3")
	id4   := c.GetString("slideshow_4")
	id5   := c.GetString("slideshow_5")
	News  := c.GetString("news")

	//用redis存储
	var client goredis.Client
	client.Addr="127.0.0.1:6379"

	err1 := client.Set("P1",[]byte(id1))
	if err1 != nil {
		panic(err1)
	}
	err2 := client.Set("P2",[]byte(id2))
	if err2 != nil {
		panic(err2)
	}
	err3 := client.Set("P3",[]byte(id3))
	if err3 != nil {
		panic(err3)
	}
	err4 := client.Set("P4",[]byte(id4))
	if err4 != nil {
		panic(err4)
	}
	err5 := client.Set("P5",[]byte(id5))
	if err5 != nil {
		panic(err5)
	}
	err6 := client.Set("P6",[]byte(News))
	if err6 != nil {
		panic(err6)
	}
	c.Redirect("/admin/home",302)
}