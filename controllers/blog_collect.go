package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
)

type BlogCollectController struct {
	JudgeController
}

func (c *BlogCollectController) Post(){

	blogId := c.GetString("collect_id")
	fmt.Println("blogId:", blogId)

	sql := fmt.Sprintf("where user_id=%d and blog_id=%s",
		models.QueryUserWithUsername(c.Loginuser.(string)), blogId)

	num := models.QueryUserCollectWithCon(sql)
	fmt.Println(num)
	//表示还未收藏这本书
	if num == 0 {
		var ct models.Collect
		ct.BlogId = blogId
		ct.UserId = strconv.Itoa(models.QueryUserWithUsername(c.Loginuser.(string)))
		_, err := models.InsertCollect(ct)

		if err != nil {
			c.responseErr(err)
			return
		}

		c.Data["json"] = map[string]interface{}{"code": 1, "message": "收藏成功"}
	}else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "收藏失败"}
	}

	c.ServeJSON()
}

func (c *BlogCollectController) responseErr(err error) {
	c.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	c.ServeJSON()
}
