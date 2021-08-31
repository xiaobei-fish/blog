package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
)

type BlogDeleteController struct {
	JudgeController
}

func (c *BlogDeleteController) Get(){
	c.Controller.EnableRender = false

	Id := c.GetString("id")
	userId := models.QueryUserWithUsername(c.Loginuser.(string))

	fmt.Printf("删除Id:%s,userId:%d\n",Id,userId)

	_, e := models.DeleteBlog(Id, strconv.Itoa(userId))
	_, E := models.DeleteTogether(Id)

	models.SetBlogRowsNum()

	if e != nil || E != nil {
		fmt.Println("删除失败")
	}else {
		fmt.Println("删除成功")
	}

	c.Redirect("/mine", 302)
}
