package controllers

import (
	"fmt"
	"web/models"
)

type ResultController struct {
	JudgeController
}

func (c *ResultController) Get()  {
	keyWords := c.GetString("keyWords")
	c.Data["Username"] = c.Loginuser
	c.Data["keyWords"] = keyWords
	userid := models.QueryUserWithUsername(c.Loginuser.(string))
	c.Data["Userid"] = userid
	c.TplName = "search.html"
}

func (c *ResultController) Post() {

	//得到要收藏书籍的信息
	userId 		 := c.GetString("userId")
	bookId 		 := c.GetString("bookId")

	fmt.Println("user:",userId)
	fmt.Println("book:",bookId)

	//检查是否已经在收藏夹中
	id := models.QueryBookWithUserId(userId,bookId)
	fmt.Println("id:", id)
	if id > 0 {
		c.Data["json"] = map[string]interface{}{"code": 0}
		c.ServeJSON()
		return
	}

	//实例化model，将它存入到数据库中
	m := models.Collect{UserId: userId,BookId:bookId}
	_, err := models.CollectNovel(m)
	_, _ = models.UpdateRankByCollectNum(bookId)

	//返回数据给浏览器
	if err == nil {
		//无误
		c.Data["json"] = map[string]interface{}{"code": 1}   //用code给前端来判断收藏成不成功
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0}
	}
	c.ServeJSON()
}
