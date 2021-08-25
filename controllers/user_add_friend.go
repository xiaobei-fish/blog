package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
)

type AddController struct {
	JudgeController
}

func (c *AddController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
		c.Data["Id"] = models.QueryUserWithUsername(c.Loginuser.(string))
	}

	c.TplName = "add.html"
}

func (c *AddController) Post() {
	myId := models.QueryUserWithUsername(c.Loginuser.(string))
	f_phone := c.GetString("phone")

	fmt.Printf("myId:%s,friend_phone:%s\n",myId,f_phone)
	f_Id := models.QueryUserWithPhone(f_phone)

	id := models.QueryUserWithPhoneAndId(f_phone, strconv.Itoa(myId))

	if id > 0 {
		//电话是自己的号码
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败,所填号码是自己的"}
	}else {
		_, err := models.InsertFriend(strconv.Itoa(myId), strconv.Itoa(f_Id))
		if err != nil {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}
		} else {
			c.Data["json"] = map[string]interface{


			}{"code": 1, "message": "添加成功"}
		}
	}
	c.ServeJSON()
}
