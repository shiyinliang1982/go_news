package routers

import (
	"github.com/astaxie/beego"
	"news/controllers"
)

func init() {
	//注册功能
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:PostRegister")

	//登陆功能
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:PostLogin")

	//文章列表页面访问
	beego.Router("/article", &controllers.ArticleController{}, "get:ShowArticleList")

	//添加文章
	beego.Router("/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:AddArticle")

	//显示文章详情
	beego.Router("/showArticleDetail", &controllers.ArticleController{}, "get:ShowArticleDetail")

	//编辑功能
	beego.Router("/updateArticle", &controllers.ArticleController{}, "get:UpdateArticle;post:HandleUpdateArticle")

	//删除文章功能
	beego.Router("/deleteArticle", &controllers.ArticleController{}, "get:DeleteArticle")
}
