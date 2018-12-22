package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"news/models"
)

type UserController struct {
	beego.Controller
}

//显示注册页面
func (c *UserController) ShowRegister() {
	c.TplName = "register.html"
}

//注册功能
func (c *UserController) PostRegister() {
	//1.获取数据
	userName := c.GetString("userName")
	passWord := c.GetString("passWord")

	//2.校验数据
	if userName == "" || passWord == "" {
		c.Data["errmsg"] = "用户名或密码不能为空"
		c.TplName = "register.html"
		return
	}
	//3.操作数据
	//3.1获取orm对象
	o := orm.NewOrm()

	//3.2创建操作对象
	var user models.User
	user.UserName = userName
	user.PassWord = passWord

	//3.3通过操作对象操作数据库
	_, err := o.Insert(&user)
	if err != nil {
		c.Data["errmsg"] = "用户注册失败"
		c.TplName = "register.html"
		return
	}

	//4.注册成功,重定向到登陆页面
	c.Redirect("/login", 302)
}

//显示登陆页面
func (c *UserController) ShowLogin() {
	c.TplName = "login.html"
}

func (c *UserController) PostLogin() {
	//1.获取数据
	userName := c.GetString("userName")
	passWord := c.GetString("passWord")

	//2.校验数据
	if userName == "" || passWord == "" {
		c.Data["errmsg"] = "用户名或密码不能为空"
		c.TplName = "login.html"
		return
	}

	//3.操作数据
	//3.1获取orm对象
	o := orm.NewOrm()

	//3.2获取操作数据库对象
	var user models.User
	user.UserName = userName

	//3.3操作数据库
	err := o.Read(&user, "UserName")
	if err != nil {
		c.Data["errmsg"] = "用户名不存在"
		c.TplName = "login.html"
		return
	}
	if user.PassWord != passWord {
		c.Data["errmsg"] = "密码错误"
		c.TplName = "login.html"
		return
	}
	//4.返回页面
	//c.Ctx.WriteString("登陆成功")
	c.Redirect("/article", 302)
}
