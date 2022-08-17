package controllers

type AdminExitController struct {
	JudgeController
}

func (c *AdminExitController)Get(){
	//清除该用户登录状态的数据
	c.DelSession("loginuser")
	c.Redirect("/home",302)
}

