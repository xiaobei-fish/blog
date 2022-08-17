package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})								//首页
	beego.Router("/register", &controllers.RegisterController{})					//注册
	beego.Router("/login", &controllers.LoginController{})						//登录
	beego.Router("/exit",&controllers.ExitController{})							//退出
	beego.Router("/alarm",&controllers.AlarmController{})							//提示
	//beego.Router("/forget",&controllers.ForgetController{})						//忘记密码
	beego.Router("/selfhome",&controllers.SelfHomeController{})					//个人页面
	beego.Router("/img",&controllers.ImgController{})								//头像更新
	beego.Router("/alter",&controllers.AlterController{})							//修改密码
	beego.Router("/search",&controllers.SearchController{})						//搜索界面
	beego.Router("/all",&controllers.AllBlogController{})							//全部博客
	beego.Router("/write",&controllers.WriteController{})							//撰写博客
	beego.Router("/mine",&controllers.MineController{})							//我的博客
	beego.Router("/collect",&controllers.CollectController{})						//我的收藏
	beego.Router("/detail",&controllers.DetailController{})						//博客详细
	beego.Router("/like",&controllers.LikeController{})							//点赞博客
	beego.Router("/blog_collect",&controllers.BlogCollectController{})			//收藏博客
	beego.Router("/delete",&controllers.BlogDeleteController{})					//删除博客
	beego.Router("/update",&controllers.UpdateBlogController{})					//更新博客
	beego.Router("/delete_collect",&controllers.CollectDeleteController{})		//删除收藏
	beego.Router("/friend",&controllers.FriendController{})						//我的好友
	beego.Router("/add",&controllers.AddController{})								//添加好友
	beego.Router("/info_center",&controllers.CenterController{})					//好友界面
	beego.Router("/note",&controllers.NoteController{})							//留言页面
}
