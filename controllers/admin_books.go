package controllers

import (
	"web/models"
)

type AdminBooksController struct {
	JudgeController
}

func (c *AdminBooksController) Get() {
	c.Data["adminName"] = c.Loginuser
	c.TplName = "admin_books.html"
}

func (c *AdminBooksController) Post()  {
	id := c.GetString("change_id")
	name := c.GetString("change_name")
	state := c.GetString("change_state")
	genre := c.GetString("change_genre")
	introduction := c.GetString("change_introduction")
	/*fmt.Println("id:",id)
	fmt.Println("name:",name)
	fmt.Println("state:",state)
	fmt.Println("genre:",genre)
	fmt.Println("introduction:",introduction)*/
	models.UpdateNovel(id, name, state, genre, introduction)
	c.TplName = "admin_books.html"
}