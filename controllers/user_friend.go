package controllers

import (
	"blog/models"
)

type FriendController struct {
	JudgeController
}

func (c *FriendController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
		c.Data["Img"] = models.QueryUserHeadImgWithUsername(c.Loginuser.(string))
	}
	friendList := models.QueryUserFriendsWithId(models.QueryUserWithUsername(c.Loginuser.(string)))
	c.Data["Content"] = models.MakeFriendBlocks(friendList)

	c.TplName = "friend.html"
}
