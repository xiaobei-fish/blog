package controllers

import (
	"fmt"
	"log"
	"strconv"
	"web/models"
)

type PageController struct {
	JudgeController
}

func (c *PageController) Get() {
	c.Data["Username"] = c.Loginuser
	userid := models.QueryUserWithUsername(c.Loginuser.(string))
	c.Data["Userid"] = userid
	c.TplName = "collect.html"
}

func (c *PageController) Post() {
	//得到小说块小说对应的id
	novelID := c.GetString("id")
	userid := models.QueryUserWithUsername(c.Loginuser.(string))
	userId := strconv.Itoa(userid)

	fmt.Println("删除的收藏小说id:", novelID)

	_, err := models.DeleteCollectedNovel(novelID,userId)
	_, _ = models.UpdateRankByDelete(novelID)
	if err != nil {
		log.Println(err)
	}
	c.Redirect("/page", 302)
}
