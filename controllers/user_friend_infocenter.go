package controllers

import "blog/models"

type CenterController struct {
	JudgeController
}

func (c *CenterController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
	}
	friendId := c.GetString("id")
	friendName := models.QueryUsernameWithId(friendId)
	c.Data["Img"] = models.QueryUserHeadImgWithUsername(friendName)

	uList := models.QueryUserInfoWithUsername(friendName)
	for _, u := range uList {
		c.Data["userId"] = u.Id
		c.Data["Phone"] = u.Phone
		c.Data["username"] = u.Username
	}
	noteList := models.QueryNoteWithNotedId(friendId)
	c.Data["Content"] = models.MakeSmallNoteBlocks(noteList)

	c.TplName = "info_center.html"
}
