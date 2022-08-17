package controllers

type ExitController struct {
	JudgeController
}

func (c *ExitController)Get(){
	//清除该用户登录状态的数据
	c.DelSession("loginuser")
	c.Redirect("/home",302)
}
