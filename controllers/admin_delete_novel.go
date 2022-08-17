package controllers

import (
	"fmt"
	"log"
	"strconv"
	"web/models"
)


type DeleteNovelController struct {
	JudgeController
}

//点击删除后重定向到管理员书籍页
func (c *DeleteNovelController) Post() {
	novelID := c.GetString("remove_id")//得到要删除小说对应的id
	fmt.Println("删除id:", novelID)
	id,_ := strconv.Atoi(novelID)

	_, err := models.DeleteNovel(id)
	if err != nil {
		log.Println(err)
	}
	c.Redirect("/admin/books", 302)
}

