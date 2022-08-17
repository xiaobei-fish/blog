package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func init(){
	beego.Router("/home",&controllers.HomeController{})						//首页
	beego.Router("/selfhome",&controllers.SelfhomeController{})				//个人首页
	beego.Router("/register", &controllers.RegisterController{})				//注册
	beego.Router("/login", &controllers.LoginController{})					//登录
	beego.Router("/exit",&controllers.ExitController{})						//退出
	beego.Router("/alter",&controllers.UserAlterPasswordController{})			//修改密码
	beego.Router("/book/:id",&controllers.BookDetailController{})				//小说详细页

	beego.Router("/admin", &controllers.AdminController{})					//管理员登录
	beego.Router("/admin/home", &controllers.AdminHomeController{})			//管理员首页
	beego.Router("/admin/books", &controllers.AdminBooksController{})			//管理员管理书籍
	beego.Router("/admin/books/delete", &controllers.DeleteNovelController{}) //管理员删除书籍
	beego.Router("/admin/user/delete", &controllers.DeleteUserController{})	//管理员删除书籍
	beego.Router("/admin/user", &controllers.AdminUserController{})			//管理员管理用户
	beego.Router("/admin/exit",&controllers.AdminExitController{})			//管理员退出

	beego.Router("/result",&controllers.ResultController{})					//渲染搜索结果展示URL,以及用户收藏
	beego.Router("/search",&controllers.SearchController{})					//搜索结果jsonURL
	beego.Router("/writer",&controllers.WriterController{})					//作家页面
	beego.Router("/create/?page=:page",&controllers.UrlController{})			//作家小说jsonURL

	beego.Router("/page",&controllers.PageController{})						//书库数据渲染URL
	beego.Router("/collect/?page=:page",&controllers.CollectController{})		//书库数据jsonURL
	beego.Router("/admin_books/?page=:page",&controllers.Url1Controller{})	//管理员书库数据jsonURL
	beego.Router("/admin_user/?page=:page",&controllers.Url2Controller{}) 	//管理员用户数据jsonURL

	//爬虫专用
	beego.Router("/crawl_novel", &controllers.CrawlNovelController{},"*:CrawlNovel")
}
