package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
)

type CollectDeleteController struct {
	JudgeController
}

func (c *CollectDeleteController) Get(){
	c.Controller.EnableRender = false

	Id := c.GetString("id")
	userId := models.QueryUserWithUsername(c.Loginuser.(string))

	fmt.Printf("删除Id:%s,userId:%d\n",Id,userId)

	_, _ = models.DeleteCollect(strconv.Itoa(userId), Id)

	fmt.Println("删除成功")

	c.Redirect("/collect", 302)
}
