package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"news/models"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}

//显示文章列表页面
func (ac *ArticleController) ShowArticleList() {
	//1获取数据
	//1.1指定表
	o := orm.NewOrm()
	/*高级查询
	* 参数是表名
	* 返回值是查询集
	 */
	querySeter := o.QueryTable("Article")
	var articles []models.Article
	/*_, err := querySeter.All(&articles)
	if err != nil {
		beego.Error("查询数据错误",err)
	}*/

	/*查询总记录数
	* 返回值1：count总记录数
	* 返回值2：操作错误
	 */
	count, err := querySeter.Count()
	if err != nil {
		beego.Error("查询数据错误", err)
	}

	//查询总页数
	pageSize := 5
	pageCount := math.Ceil(float64(count) / float64(pageSize))

	//获取页码
	pageIndex, err := ac.GetInt("pageIndex")
	if err != nil {
		pageIndex = 1
	}

	/*根据页码获取部分数据
	* 第一个参数获取几条数据
	* 第二个参数从哪条数据开始获取,需要动态获取
	 */
	start := (pageIndex - 1) * pageSize
	querySeter.Limit(pageSize, start).All(&articles)

	//2传递数据
	ac.Data["pageIndex"] = pageIndex
	ac.Data["pageCount"] = int(pageCount)
	ac.Data["count"] = count
	ac.Data["articles"] = articles
	ac.TplName = "index.html"
}

//显示添加文章页面
func (ac *ArticleController) ShowAddArticle() {
	ac.TplName = "add.html"
}

//获取添加文章数据
func (ac *ArticleController) AddArticle() {
	//1.获取数据
	articleName := ac.GetString("articleName")
	content := ac.GetString("content")

	//1.1处理上传文件
	fileName := UploadFile(&ac.Controller, "uploadname", "add.html")

	//2.校验数据
	if articleName == "" || content == "" || fileName == "" {
		ac.Data["errmsg"] = "文章标题或内容不完整"
		ac.TplName = "add.html"
		return
	}

	//3.处理数据
	o := orm.NewOrm()

	var article models.Article
	article.ArtiName = articleName
	article.Acontent = content
	article.Aimg = fileName

	_, err := o.Insert(&article)
	if err != nil {
		ac.Data["errmsg"] = "文章保存失败"
		ac.TplName = "add.html"
		return
	}

	//4.返回页面
	ac.Redirect("/article", 302)
}

//展示文章详情页面
func (ac *ArticleController) ShowArticleDetail() {
	//1.获取数据
	articleId, err := ac.GetInt("articleId")

	//2.校验数据
	if err != nil {
		beego.Error("传递链接错误", err)
	}

	//3.处理数据
	o := orm.NewOrm()

	var article models.Article
	article.Id = articleId

	err = o.Read(&article)
	if err != nil {
		beego.Info("查询失败", err)
	}

	//修改阅读量
	article.Acount += 1
	_, err = o.Update(&article)
	if err != nil {
		beego.Info("修改阅读量失败", err)
	}

	//4.返回页面
	ac.Data["article"] = article
	ac.TplName = "content.html"
}

//显示编辑功能页面
func (ac *ArticleController) UpdateArticle() {
	//1.获取数据
	articleId, err := ac.GetInt("articleId")

	//2.校验数据
	if err != nil {
		beego.Error("获取文章id失败", err)
		return
	}

	//3.处理数据
	o := orm.NewOrm()

	var article models.Article
	article.Id = articleId

	err = o.Read(&article)
	if err != nil {
		beego.Info("查询失败", err)
	}

	//4.返回页面
	ac.Data["article"] = article
	ac.TplName = "update.html"
}

/*封装文件上传功能
* 第一个参数:控制器 类型为控制器的父类
* 第二个参数:从页面获取的上传文件名 类型为字符串
* 第三个参数:发生错误时返回页面的视图文件
* 返回值:存储文件的路径
 */
func UploadFile(ac *beego.Controller, filePath, view string) string {
	/*获取上传图片
	 * 参数为前端name属性的值
	 * 返回值1:file文件的字节流
	 * 返回值2:header响应头
	 * 返回值3:err获取上传文件错误
	 */
	file, header, err := ac.GetFile(filePath)

	if err != nil {
		ac.Data["errmsg"] = "图片上传失败"
		ac.TplName = view
		return ""
	}

	defer file.Close()
	//2.1文件大小
	if header.Size > 5000000 {
		ac.Data["errmsg"] = "文件太大，请重新上传"
		ac.TplName = view
		return ""
	}

	//2.2文件格式
	ext := path.Ext(header.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		ac.Data["errmsg"] = "文件格式错误，请重新上传"
		ac.TplName = view
		return ""
	}

	/*2.3防止重名
	 * 参数:"2006-01-02-15:04:05"为固定格式,不能更改
	 */
	fileName := time.Now().Format("2006-01-02-15:04:05") + ext

	/*存储文件
	* 参数1：前端name属性的值
	* 参数2：文件存储路径，路径使用相对路径
	 */
	err = ac.SaveToFile(filePath, "./static/img/"+fileName)
	if err != nil {
		ac.Data["errmsg"] = "图片保存失败"
		ac.TplName = view
		return ""
	}
	return "/static/img/" + fileName
}

//实现编辑文章功能
func (ac *ArticleController) HandleUpdateArticle() {
	//获取数据
	articleId, err := ac.GetInt("articleId")
	articleName := ac.GetString("articleName")
	content := ac.GetString("content")
	fileName := UploadFile(&ac.Controller, "uploadname", "update.html")

	//校验数据
	if err != nil {
		ac.Data["errmsg"] = "文章Id获取失败"
		ac.TplName = "update.html"
		return
	}

	if articleName == "" || content == "" || fileName == "" {
		ac.Data["errmsg"] = "更新文章内容不完整"
		ac.TplName = "update.html"
		return
	}

	//处理数据
	o := orm.NewOrm()
	var article models.Article
	article.Id = articleId

	err = o.Read(&article)
	if err != nil {
		ac.Data["errmsg"] = "您所修改的文章不存在"
		ac.TplName = "update.html"
		return
	}

	article.ArtiName = articleName
	article.Acontent = content
	article.Aimg = fileName
	_, err = o.Update(&article)
	if err != nil {
		ac.Data["errmsg"] = "文章更新失败"
		ac.TplName = "update.html"
		return
	}

	//返回视图
	ac.Redirect("/article", 302)
}

//删除文章功能
func (ac *ArticleController) DeleteArticle() {
	//1.获取数据
	articleId, err := ac.GetInt("articleId")

	//2.校验数据
	if err != nil {
		beego.Error("获取Id失败", err)
		return
	}

	//3.处理数据
	o := orm.NewOrm()
	var article models.Article
	article.Id = articleId
	_, err = o.Delete(&article)
	if err != nil {
		beego.Error("删除文章失败", err)
		return
	}

	//4.返回视图
	ac.Redirect("/article", 302)
}
