package controllers

type AdminUserController struct {
	JudgeController
}

func (c *AdminUserController) Get() {
	c.Data["adminName"] = c.Loginuser
	c.TplName = "admin_user.html"
}
