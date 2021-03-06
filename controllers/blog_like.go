package controllers

import (
	"blog/models"
	"fmt"
)

type LikeController struct {
	JudgeController
}

func (c *LikeController) Post() {
	blogId := c.GetString("blog_id")
	fmt.Println("blogId:", blogId)

	_,err := models.UpdateLike(blogId)

	if err != nil {
		c.responseErr(err)
		return
	}

	c.Data["json"] = map[string]interface{}{"code": 1, "message": "ηΉθ΅ζε"}

	c.ServeJSON()
}

func (c *LikeController) responseErr(err error) {
	c.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	c.ServeJSON()
}