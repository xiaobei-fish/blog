package controllers

type AlarmController struct {
	JudgeController
}

func (c *AlarmController) Get() {
	if c.Loginuser != nil {
		c.Data["Username"] = c.Loginuser.(string)
	}
	flag := c.GetString("success")
	if flag == "1"{
		c.Data["Success"] = true
	}else {
		c.Data["Success"] = false
	}
	page,_ := c.GetInt("page")
	c.Data["Page"] = page
	c.Data["Key"] = c.GetString("key")

	c.TplName = "alarm.html"
}
