package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
)

type WriteController struct {
	JudgeController
}

func (c *WriteController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
	}
	c.TplName = "write.html"
}

func (c *WriteController) Post() {
	//获取浏览器传输的数据，通过表单的name属性获取值
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中
	blog := models.Blog{UserId: strconv.Itoa(models.QueryUserWithUsername(c.Loginuser.(string))),
		Title: title, Tag: tags, Short: short, Content: content, Author: c.Loginuser.(string), Like: "0", Read: 0}
	_, err := models.InsertBlog(blog)

	//返回数据给浏览器
	var response map[string]interface{}

	if err == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}
	c.Data["json"] = response

	c.ServeJSON()
}
