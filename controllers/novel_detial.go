package controllers

import (
	"fmt"
	"web/models"
)

type BookDetailController struct {
	JudgeController
}

func (c *BookDetailController) Get() {

	idStr := c.Ctx.Input.Param(":id")

	fmt.Println("id::::::::",idStr)

	detail := models.FindNovelWithId(idStr)

	c.Data["id"]  		   = detail[0]["id"].(string)
	c.Data["picture"] 	   = detail[0]["novel_pic"].(string)
	c.Data["name"]		   = detail[0]["novel_name"].(string)
	c.Data["genre"] 	   = detail[0]["novel_genre"].(string)
	c.Data["state"]        = detail[0]["novel_state"].(string)
	c.Data["writer"] 	   = detail[0]["novel_writer"].(string)
	c.Data["introduction"] = detail[0]["novel_introduction"].(string)

	c.Data["Username"] = c.Loginuser
	c.Data["Userid"]   = models.QueryUserWithUsername(c.Loginuser.(string))
	c.TplName = "books.html"
}
